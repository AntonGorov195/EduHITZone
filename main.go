package main

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

//
// func main() {
// 	pages.AddIndexHandles()
// 	fs := http.FileServer(http.Dir("public/"))
// 	http.Handle("/", fs)
//
// 	pages.AddLoginHandles()
// 	http.ListenAndServe(":42069", nil)
// }

// Working on a single page app.
func main() {
	
	http.Handle("/static/", http.FileServer(http.Dir("SPAPublic/")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Init buffer
		var buf bytes.Buffer
		defer func() {
			// If it wasn't caused by htmx, render the whole page.
			if r.Header.Get("HX-Request") == "" {
				tmpl, _ := template.ParseFiles("SPAPublic/index.html")
				tmpl.Execute(w, template.HTML(buf.String()))
				return
			}
			tmpl, _ := template.New("view").Parse("{{ . }}")
			tmpl.Execute(w, template.HTML(buf.String()))
		}()
		// Get view
		switch strings.Trim(r.URL.Path, "/ ") {
		case "":
			{
				tmpl, err := template.ParseFiles("SPAPublic/static/views/entry.html")
				if err != nil {
					panic(err)
				}
				if err := tmpl.Execute(&buf, nil); err != nil {
					panic(err)
				}
				return
			}
		case "login":
			{
				tmpl, err := template.ParseFiles("SPAPublic/static/views/login.html")
				if err != nil {
					panic(err)
				}
				if err := tmpl.Execute(&buf, nil); err != nil {
					panic(err)
				}
				return
			}
		case "content":
			{
				tmpl, err := template.ParseFiles("SPAPublic/static/views/content.html")
				if err != nil {
					panic(err)
				}
				if err := tmpl.Execute(&buf, nil); err != nil {
					panic(err)
				}
				return
			}
		default:
			{
				tmpl, _ := template.ParseFiles("SPAPublic/error.html")
				tmpl.Execute(&buf, nil)
				return
			}
		}
	})
	http.HandleFunc("/api/v1/login/form", func(w http.ResponseWriter, r *http.Request) {
		// Contains the data needed for the error message element.
		type ErrorMessageData struct {
			Name string
		}

		r.ParseForm()
		const success = "<span id=\"error-msg\" style=\"color:green;\">Hello {{ .Name }}! Later this will look into the database and check password</span>"
		const failed = "<span id=\"error-msg\" style=\"color:red;\">Try again, {{ .Name }}! Your name is not cool</span>"

		name := r.Form.Get("name")
		if name == "Ofek" || name == "Anton" {
			tmpl, err := template.New("cool_name").Parse(success)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(w, ErrorMessageData{Name: name})
			if err != nil {
				panic(err)
			}
			return
		}
		tmpl, err := template.New("invalid_name").Parse(failed)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, ErrorMessageData{Name: name})
		if err != nil {
			panic(err)
		}

	})

	http.ListenAndServe(":42069", nil)
}
