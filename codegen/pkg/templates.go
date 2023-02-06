package pkg

import "text/template"

func constTemplate() *template.Template {
	return template.Must(template.New("constTpl").Parse("\t{{.Name}}{{.Value}}\n"))
}

func typeTemplate() *template.Template {
	return template.Must(template.New("typeTpl").Parse("type {{.Name}} {{.TypeValue}}\n"))
}

func structTemplate() *template.Template {
	return template.Must(template.New("structTpl").Parse(
		"type {{.Name}} struct {\n{{range .Types}}\t{{.Name}} {{if .TypeValue}} *{{ .TypeValue }} {{end}}{{if .TypeValues}} []*{{ .TypeValues }} {{end}} {{.Tag}}\n{{end}}}\n",
	))
}

func modelTemplate() *template.Template {
	return template.Must(template.New("modelTpl").Parse(
		"class {{.Name}}(BaseModel):\n{{range .Types}}    {{.Name}}: {{if .TypeValue}} {{ .TypeValue }} {{end}}{{if .TypeValues}} list[{{ .TypeValues}}] {{end}}\n{{end}}\n",
	))
}

func structConstructorTemplate() *template.Template {
	return template.Must(template.New("structTpl").Parse(
		"func new{{.Name}}() *{{.Name}} {\n\treturn &{{.Name}}{}\n}\n",
	))
}

func switcherTemplate() *template.Template {
	return template.Must(template.New("switcherTpl").Parse(
		"func {{.Signature}} {\n\tswitch value {\n{{range .Cases}}\t\tcase {{.CaseValue}}:\n{{if .CaseBody}}\t\t\t{{.CaseBody}}{{end}}\t\t\treturn {{.ReturnValue}}\n{{end}}\t\tdefault:\n\t\t\treturn {{.DefaultReturnValue}}\n\t}\n}\n",
	))
}
