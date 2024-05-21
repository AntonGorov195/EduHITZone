package spa

import (
	"html/template"
	"net/http"
)

func addNewAccountHandle() {
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

	})

}
