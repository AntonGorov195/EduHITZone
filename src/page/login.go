package spa

import (
	"database/sql"
	"html/template"
	"net/http"
)

func addLoginHandle(db *sql.DB) {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "login")
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
