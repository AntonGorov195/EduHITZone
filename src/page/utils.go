package spa

import (
	"bytes"
	"html/template"
	"net/http"
)

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
