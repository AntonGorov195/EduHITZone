package main

import (
	_ "database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name string
}

func main() {

	// http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl, _ := template.ParseFiles("public/index.html")
	// 	tmpl.Execute(w, nil)
	// })

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

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

	http.ListenAndServe(":42069", nil)
}
