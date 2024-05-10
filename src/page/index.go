package page

import (
	"bytes"
	"html/template"
	"net/http"
)

func AddIndexHandles() {
	http.HandleFunc("/index.html", registerIndexPage)
	// When you add a "/" pattern it will refer to all the subfolders and elements.
	// "{$}" Means it will not look into the subresourses.
	// This is needed because we need to apply the template to index page.
	// However we already using the "/" pattern to make public folder the root
	// of the file system.
	http.HandleFunc("/{$}", registerIndexPage)
}
func registerIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl_main, _ := template.ParseFiles("public/index.html")
	tmpl, _ := template.ParseFiles("public/template.html")

	// Reads the main content of the page.
	var buf bytes.Buffer
	if err := tmpl_main.Execute(&buf, nil); err != nil {
		panic("Error handling not implemented.")
	}
	page_main_content := buf.String()

	// Add the main content to the template, which will be rendered.
	tmpl.ExecuteTemplate(w, "primary-v1", PageTemplateData{
		[]template.HTML{
			GetTitleElement(tmpl, "EduHITZone"),
			GetCSSElement(tmpl, "common.css"),
			GetCSSElement(tmpl, "style.css"),
		},
		template.HTML(page_main_content),
		"",
	})
}
