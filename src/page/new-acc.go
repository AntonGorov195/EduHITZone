package spa

import (
	hitdb "EduHITZone/src/MySQL"
	"database/sql"
	"html/template"
	"net/http"
)

func addNewAccountHandle(db *sql.DB) {
	http.HandleFunc("/new-acc", func(w http.ResponseWriter, r *http.Request) {
		DrawView(w, r, "new-acc")
	})
	http.HandleFunc("/api/v1/new-acc/form", func(w http.ResponseWriter, r *http.Request) {
		type ErrorMessageData struct {
			Name string
		}

		r.ParseForm()
		const success = "<span id=\"error-msg\" style=\"color:red;\">The name \"{{ .Name }}\" is already taken. Use a different one.</span>"
		const failed = "<span id=\"error-msg\" style=\"color:green;\">Good, now I have to add it to the database. I will need to talk with Ofek Biton about it.</span>"

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
		hitdb.AddStudent(db, r.Form.Get("name"), r.Form.Get("name"), r.Form.Get("email"), 1, r.Form.Get("date-of-birth"))
	})

}
