package spa

import (
	"bytes"
	"html/template"
	"net/http"
)

// Write the view into the writer. If this wasn't an htmx request then it will assume a complete page re-render.
// This is useful with hx-push-url. If it isn't HTMX request then the user entered the url manually, meaning the site was reloaded.
// However, if it was an HTMX request then the site is already loaded, meaning you only need load the view.
func sendViewBuf(w http.ResponseWriter, r *http.Request, view bytes.Buffer) {
	if r.Header.Get("HX-Request") == "" {
		// If not HTMX, re-render the page.
		tmpl, err := template.ParseFiles("public/index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = tmpl.Execute(w, template.HTML(view.String()))

		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}
	// If it is HTMX then render the view.
	tmpl, err := template.New("view").Parse("{{ . }}")
	if err != nil {
		http.Error(w, "Something went really wrong. This must mean that the Go std no longer functions.", 500)
		return
	}

	err = tmpl.Execute(w, template.HTML(view.String()))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Loads view into a buffer.
func loadView(view_name string, data any) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tmpl, err := template.ParseFiles("public/static/views/" + view_name + ".html")
	if err != nil {
		// Return empty buffer
		return bytes.Buffer{}, err
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		// Return empty buffer
		return bytes.Buffer{}, err
	}
	return buf, nil
}

// Default way to draw a view.
func drawView(w http.ResponseWriter, r *http.Request, view_name string) {
	// Parse the view.
	view, err := loadView(view_name, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	sendViewBuf(w, r, view)
}
