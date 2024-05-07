package loginPage

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Name string
}

func AddHandles() {
	http.HandleFunc("/login/login.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, world!")
		tmpl, _ := template.ParseFiles("public/login/login.html")
		tmpl.Execute(w, 1)
	})

	http.HandleFunc("/tmp/1", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.Form.Get("name")
		if name == "Ofek" || name == "Anton" {
			tmpl, err := template.New("tmp").Parse("Hello {{.Name}}! Later this will look into the database and check password")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(w, Data{Name: r.Form.Get("name")})
			if err != nil {
				panic(err)
			}
			return
		}
		tmpl, err := template.New("tmp").Parse("Try again, {{.Name}}! Your name is not cool :(")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, Data{Name: r.Form.Get("name")})
		if err != nil {
			panic(err)
		}
	})
}
