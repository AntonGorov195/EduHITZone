package loginPage

import (
	"fmt"
	"html/template"
	"net/http"
)

func AddHandles() {
	// The initial data for the login page.
	type LoginPagInitData struct {
		ErrId string
	}
	// The ID of the element that will contain the message send by the form.
	const ErrElemId = "error-msg"
	http.HandleFunc("/login/login.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, world!")
		tmpl, _ := template.ParseFiles("public/login/login.html")

		tmpl.Execute(w, LoginPagInitData{ErrElemId})
	})

	http.HandleFunc("/api/v1/login/form", func(w http.ResponseWriter, r *http.Request) {
		// Contains the data needed for the error message element.
		type ErrorMessageData struct {
			Name string
		}

		r.ParseForm()
		const success = "<span id=\"" + ErrElemId + "\" style=\"color:green;\">Hello {{ .Name }}! Later this will look into the database and check password</span>"
		const failed = "<span id=\"" + ErrElemId + "\" style=\"color:red;\">Try again, {{ .Name }}! Your name is not cool</span>"

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
