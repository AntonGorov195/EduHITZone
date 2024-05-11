package page

import (
	"bytes"
	"html/template"
	"net/http"
)

// The ID of the element that will contain the message send by the form.
const errElemId = "error-msg"

func AddLoginHandles() {
	http.HandleFunc("/login/login.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl_main, _ := template.ParseFiles("public/login/login.html")
		tmpl, _ := template.ParseFiles("public/template.html")

		// Data for the main part of the login page.
		// Only needs the Id of the element where the error message would be displayed.
		data := struct {
			ErrId string
		}{
			errElemId,
		}

		var buf bytes.Buffer
		// Reads the main content of the page.
		if err := tmpl_main.Execute(&buf, data); err != nil {
			panic(err)
		}
		page_main_content := buf.String()

		// Add the main content to the template, which will be rendered.
		tmpl.ExecuteTemplate(w, "primary-v1", PageTemplateData{
			[]template.HTML{
				GetTitleElement(tmpl, "Log Into EduHITZone"),
				GetCSSElement(tmpl, "../common.css"),
				GetCSSElement(tmpl, "login.css"),
			},
			template.HTML(page_main_content),
			"Login"})
	})
	http.HandleFunc("/api/v1/login/form", func(w http.ResponseWriter, r *http.Request) {
		// Contains the data needed for the error message element.
		type ErrorMessageData struct {
			Name string
		}

		r.ParseForm()
		const success = "<span id=\"" + errElemId + "\" style=\"color:green;\">Hello {{ .Name }}! Later this will look into the database and check password</span>"
		const failed = "<span id=\"" + errElemId + "\" style=\"color:red;\">Try again, {{ .Name }}! Your name is not cool</span>"

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
}
