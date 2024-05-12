package spa

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func AddLoginHandle() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		tmpl, err := template.ParseFiles("SPAPublic/static/views/login.html")
		if err != nil {
			fmt.Println("Failed to parse login.html template")
			panic(err.Error())
		}
		if err := tmpl.Execute(&buf, nil); err != nil {
			fmt.Println("Failed to execute login.html template")
			panic(err.Error())
		}
		Render(w, r, buf)
	})
	http.HandleFunc("/api/v1/login/form", func(w http.ResponseWriter, r *http.Request) {
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
}
