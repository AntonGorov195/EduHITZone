package page

import (
	"bytes"
	"html/template"
)

type PageTemplateData struct {
	HeadEx   []template.HTML
	Main     template.HTML
	PageName string
}

func GetCSSElement(tmpl *template.Template, path string) template.HTML {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "css", path); err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

func GetTitleElement(tmpl *template.Template, title string) template.HTML {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "title", title); err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}
