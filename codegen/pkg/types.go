package pkg

type XMLElement struct {
	name     string
	level    int
	id       string
	multiple bool
}

type NameTemplate struct {
	Name string
}

type ConstTemplate struct {
	NameTemplate
	Value string
}

type BaseTypeTemplate struct {
	NameTemplate
	TypeValue  string
	TypeValues string
}

type ModelTemplate struct {
	NameTemplate
	Types []BaseTypeTemplate
}

type TypeTemplate struct {
	BaseTypeTemplate
	Tag            string
	HasConstructor bool
}

type StructTemplate struct {
	NameTemplate
	Types []TypeTemplate
}

type FuncTemplate struct {
	Signature string
}

type SwitchCaseTemplate struct {
	CaseValue   string
	CaseBody    string
	ReturnValue string
}

type SwitchTemplate struct {
	FuncTemplate
	Cases              []SwitchCaseTemplate
	DefaultReturnValue string
}

func NewNameTemplate(name string) NameTemplate {
	return NameTemplate{Name: name}
}

func NewBaseTypeTemplate(name string, typeValue string, typeValues string) BaseTypeTemplate {
	return BaseTypeTemplate{
		NameTemplate: NewNameTemplate(name),
		TypeValue:    typeValue,
		TypeValues:   typeValues,
	}
}

func NewConstTemplate(name string, value string) ConstTemplate {
	return ConstTemplate{NewNameTemplate(name), value}
}

func NewTypeTemplate(name string, tag string, typeValue string, typeValues string, hasConstructor bool) TypeTemplate {
	return TypeTemplate{NewBaseTypeTemplate(name, typeValue, typeValues), tag, hasConstructor}
}

func NewSwitchCaseTemplate(caseValue string, caseBody string, returnValue string) SwitchCaseTemplate {
	return SwitchCaseTemplate{caseValue, caseBody, returnValue}
}

func NewFuncTemplate(signature string) FuncTemplate {
	return FuncTemplate{signature}
}

func NewSwitchTemplate(signature string, cases []SwitchCaseTemplate, defaultReturnValue string) SwitchTemplate {
	return SwitchTemplate{NewFuncTemplate(signature), cases, defaultReturnValue}
}
