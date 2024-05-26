package spa

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func DrawView(w http.ResponseWriter, r *http.Request, view_name string) {
	var buf bytes.Buffer
	tmpl, err := template.ParseFiles("SPAPublic/static/views/" + view_name + ".html")
	if err != nil {
		fmt.Println("Failed to parse " + view_name + ".html template")
		panic(err.Error())
	}
	if err := tmpl.Execute(&buf, nil); err != nil {
		fmt.Println("Failed to execute " + view_name + ".html template")
		panic(err.Error())
	}
	Render(w, r, buf)
}

func Render(w http.ResponseWriter, r *http.Request, buf bytes.Buffer) {
	// If it wasn't caused by htmx, render the whole page.
	if r.Header.Get("HX-Request") == "" {
		tmpl, _ := template.ParseFiles("SPAPublic/index.html")
		tmpl.Execute(w, template.HTML(buf.String()))
		return
	}
	tmpl, _ := template.New("view").Parse("{{ . }}")
	tmpl.Execute(w, template.HTML(buf.String()))
}
