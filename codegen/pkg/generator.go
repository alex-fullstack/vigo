package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func Generate(out *os.File, elements []XMLElement) {
	stat, _ := out.Stat()
	switch filepath.Ext(stat.Name()) {
	case ".go":
		generateGo(out, elements)
	case ".py":
		generatePy(out, elements)
	}
}

func generateGo(out *os.File, elements []XMLElement) {
	idTypeConst := make([]ConstTemplate, 0, len(elements))
	idValueConst := make([]ConstTemplate, 0, len(elements))
	idValueTypeSwitcherCases := make([]SwitchCaseTemplate, 0, len(elements))
	idTypeValueSwitcherCases := make([]SwitchCaseTemplate, 0, len(elements))
	idTypeElementSwitcherCases := make([]SwitchCaseTemplate, 0, len(elements))
	idTypeConst = append(idTypeConst, NewConstTemplate("WrongType", " EMBLIDType = iota"))
	parents := []string{"specification"}
	for i, el := range elements {
		idTypeConst = append(idTypeConst, NewConstTemplate(el.name+"IdType", ""))
		idValueConst = append(idValueConst, NewConstTemplate(el.name+"Id", ` EMBLIDValue = `+el.id))
		idValueTypeSwitcherCases = append(idValueTypeSwitcherCases, NewSwitchCaseTemplate(el.name+"Id", "", el.name+"IdType"))
		idTypeValueSwitcherCases = append(idTypeValueSwitcherCases, NewSwitchCaseTemplate(el.name+"IdType", "", el.name+"Id"))
		hasChildren := i < len(elements)-1 && elements[i+1].level > el.level
		if i > 0 && elements[i-1].level > el.level {
			parents = parents[:el.level+1]
		}
		fillTypeElementSwitcherCases(el, &idTypeElementSwitcherCases, parents, hasChildren)
		if hasChildren {
			if el.multiple {
				parents = append(parents, el.name+"[len("+strings.Join(parents, ".")+"."+el.name+")-1]")
			} else {
				parents = append(parents, el.name)
			}

		}
	}

	fmt.Fprintln(out, `package `+filepath.Base(filepath.Dir(os.Args[2])))
	fmt.Fprintln(out)
	addConstant(out, []ConstTemplate{NewConstTemplate("CRCIdValue", " = 0xbf")})
	fmt.Fprintln(out)
	addConstant(out, []ConstTemplate{NewConstTemplate("EBMLMaxIDLengthDefault", " = 4")})
	fmt.Fprintln(out)
	addConstant(out, idTypeConst)
	addConstant(out, idValueConst)
	fmt.Fprintln(out)
	addType(out, NewTypeTemplate("EMBLIDType", "", "uint", "", false))
	fmt.Fprintln(out)
	addType(out, NewTypeTemplate("EMBLIDValue", "", "int", "", false))
	fmt.Fprintln(out)
	addSpecificationStructure(out, elements, "Specification", 0)
	fmt.Fprintln(out)
	addSwitcher(
		out,
		NewSwitchTemplate(
			"convertIdValueToIdType(value EMBLIDValue) EMBLIDType",
			idValueTypeSwitcherCases,
			"WrongType",
		),
	)
	fmt.Fprintln(out)
	addSwitcher(
		out,
		NewSwitchTemplate(
			"convertIdTypeToIdValue(value EMBLIDType) EMBLIDValue",
			idTypeValueSwitcherCases,
			"-1",
		),
	)
	fmt.Fprintln(out)
	addSwitcher(
		out,
		NewSwitchTemplate(
			"getElement(specification *Specification, value EMBLIDType) (el *SpecificationElement, hasChildren bool)",
			idTypeElementSwitcherCases,
			"",
		),
	)
	fmt.Fprintln(out)
}

func generatePy(out *os.File, elements []XMLElement) {
	fmt.Fprintln(out, "from pydantic import BaseModel")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "from models import SpecificationElement")
	fmt.Fprintln(out)
	fmt.Fprintln(out)
	addSpecificationModel(out, elements, "Specification", 0)
}

func fillTypeElementSwitcherCases(el XMLElement, values *[]SwitchCaseTemplate, parents []string, hasChildren bool) {
	*values = append(
		*values,
		SwitchCaseTemplate{
			CaseValue:   el.name + "IdType",
			CaseBody:    getTypeElementSwitcherCaseBody(el, parents, hasChildren),
			ReturnValue: getTypeElementSwitcherReturnValue(el, parents, hasChildren),
		},
	)
}

func getTypeElementSwitcherCaseBody(el XMLElement, parents []string, hasChildren bool) (result string) {
	elementName := strings.Join(parents, ".") + "." + el.name
	result = "if " + elementName + " == nil {\n"
	if hasChildren {
		if el.multiple {
			result += "\t\t\t\t" + elementName + " = " + "[]*Specification" + el.name + "{newSpecification" + el.name + "()}\n"
			result += "\t\t\t} else {\n"
			result += "\t\t\t\t" + elementName + " = " + "append(" + elementName + ", newSpecification" + el.name + "())\n"
			result += "\t\t\t}\n"
		} else {
			result += "\t\t\t\t" + elementName + " = " + "newSpecification" + el.name + "()\n\t\t\t}\n"
		}
	} else {
		if el.multiple {
			result += "\t\t\t\t" + elementName + " = " + "[]*SpecificationElement{newSpecificationElement()}\n"
			result += "\t\t\t} else {\n"
			result += "\t\t\t\t" + elementName + " = " + "append(" + elementName + ", newSpecificationElement())\n"
			result += "\t\t\t}\n"
		} else {
			result += "\t\t\t\t" + elementName + " = " + "newSpecificationElement()\n\t\t\t}\n"
		}
	}
	return
}

func getTypeElementSwitcherReturnValue(el XMLElement, parents []string, hasChildren bool) (result string) {
	result = strings.Join(parents, ".") + "." + el.name
	if el.multiple {
		result += `[len(` + result + `)-1]`
	}
	if hasChildren {
		result = `&` + result + `.SpecificationElement, true`
	} else {
		result += ", false"
	}
	return
}

func addSpecificationStructure(out *os.File, elements []XMLElement, name string, level int) {
	structure := StructTemplate{NameTemplate: NameTemplate{name}, Types: make([]TypeTemplate, 0)}
	if level > 0 {
		structure.Types = append(
			structure.Types,
			NewTypeTemplate("SpecificationElement", "", "", "", false),
		)
	}
	nextLevelElements := make([]XMLElement, 0)
	nextLevelStructureName := ""
	for i, el := range elements {
		if el.level == level {
			typeValueSuffix := "Element"
			hasConstructor := false
			if i < len(elements)-1 {
				if elements[i+1].level > level {
					typeValueSuffix = el.name
					hasConstructor = true
				}
			}
			var typeTemplate TypeTemplate
			tagElName := el.name
			a := []rune(tagElName)
			a[0] = unicode.ToLower(a[0])
			tagElName = string(a)
			tag := "`json:\"" + tagElName + ",omitempty\"`"
			if el.multiple {
				typeTemplate = NewTypeTemplate(el.name, tag, "", `Specification`+typeValueSuffix, false)
			} else {
				typeTemplate = NewTypeTemplate(el.name, tag, `Specification`+typeValueSuffix, "", hasConstructor)
			}

			structure.Types = append(
				structure.Types,
				typeTemplate,
			)

			if len(nextLevelElements) != 0 {
				addSpecificationStructure(out, nextLevelElements, nextLevelStructureName, level+1)
			}
			nextLevelElements = nil
			nextLevelStructureName = `Specification` + el.name
		} else {
			nextLevelElements = append(nextLevelElements, el)
		}
	}
	if len(nextLevelElements) != 0 {
		addSpecificationStructure(out, nextLevelElements, nextLevelStructureName, level+1)
	}
	addStruct(out, structure)
	addStructConstructor(out, structure)
	fmt.Fprintln(out)
}

func addConstant(w io.Writer, data []ConstTemplate) {
	constTpl := constTemplate()
	if len(data) == 1 {
		fmt.Fprintf(w, "const ")
	} else {
		fmt.Fprintln(w, "const (")
	}
	for _, value := range data {
		constTpl.Execute(w, value)
	}
	if len(data) > 1 {
		fmt.Fprintln(w, ")")
	}
}

func addType(w io.Writer, value TypeTemplate) {
	typeTpl := typeTemplate()
	typeTpl.Execute(w, value)
}

func addStruct(w io.Writer, value StructTemplate) {
	structTpl := structTemplate()
	structTpl.Execute(w, value)
}

func addStructConstructor(w io.Writer, value StructTemplate) {
	structTpl := structConstructorTemplate()
	structTpl.Execute(w, value)
}

func addSwitcher(w io.Writer, value SwitchTemplate) {
	structTpl := switcherTemplate()
	structTpl.Execute(w, value)
}

func addModel(w io.Writer, value ModelTemplate) {
	modelTpl := modelTemplate()
	modelTpl.Execute(w, value)
}

func addSpecificationModel(out *os.File, elements []XMLElement, name string, level int) {
	model := ModelTemplate{NameTemplate: NameTemplate{Name: name}, Types: make([]BaseTypeTemplate, 0)}
	if level > 0 {
		model.Types = append(
			model.Types,
			NewBaseTypeTemplate("specification_element", "SpecificationElement", ""),
		)
	}
	nextLevelElements := make([]XMLElement, 0)
	nextLevelModelName := ""
	for i, el := range elements {
		if el.level == level {
			typeValueSuffix := "Element"
			if i < len(elements)-1 {
				if elements[i+1].level > level {
					typeValueSuffix = el.name
				}
			}
			var typeTemplate BaseTypeTemplate
			if el.multiple {
				typeTemplate = NewBaseTypeTemplate(ToSnakeCase(el.name), "", "Specification"+typeValueSuffix)
			} else {
				typeTemplate = NewBaseTypeTemplate(ToSnakeCase(el.name), "Specification"+typeValueSuffix, "")
			}

			model.Types = append(model.Types, typeTemplate)

			if len(nextLevelElements) != 0 {
				addSpecificationModel(out, nextLevelElements, nextLevelModelName, level+1)
			}
			nextLevelElements = nil
			nextLevelModelName = "Specification" + el.name
		} else {
			nextLevelElements = append(nextLevelElements, el)
		}
	}
	if len(nextLevelElements) != 0 {
		addSpecificationModel(out, nextLevelElements, nextLevelModelName, level+1)
	}
	addModel(out, model)
	fmt.Fprintln(out)
}
