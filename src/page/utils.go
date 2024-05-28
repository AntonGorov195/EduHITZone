package spa

import (
	"bytes"
	"html/template"
	"net/http"
)

// Write the view into the writer. If this wasn't an htmx request then it will assume a complete page re-render.
// This is usful with hx-push-url. If it isn't HTMX request then the user entered the url manually, meaning the site was reloaded.
// However, if it was an HTMX request then the site is already loaded, meaning you only need load the view.
func drawView(w http.ResponseWriter, r *http.Request, view_name string) {
	// Parse the view.
	var buf bytes.Buffer
	tmpl, err := template.ParseFiles("SPAPublic/static/views/" + view_name + ".html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := tmpl.Execute(&buf, nil); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if r.Header.Get("HX-Request") == "" {
		// If not HTMX, re-render the page.
		tmpl, err := template.ParseFiles("SPAPublic/index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = tmpl.Execute(w, template.HTML(buf.String()))

		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}
	// If it is HTMX then render the view.
	tmpl, err = template.New("view").Parse("{{ . }}")
	if err != nil {
		http.Error(w, "Something went really wrong. This must mean that the Go std no longer functions.", 500)
		return
	}

	err = tmpl.Execute(w, template.HTML(buf.String()))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
