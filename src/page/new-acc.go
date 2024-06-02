package spa

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	hitdb "EduHITZone/src/MySQL"
	"database/sql"
)

func addNewAccountHandle(db *sql.DB) {
	http.HandleFunc("/new-acc", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "new-acc")
	})

	http.HandleFunc("/api/v1/new-acc/form", func(w http.ResponseWriter, r *http.Request) {
		type ErrorMessageData struct {
			Name string
		}

		r.ParseForm()
		const success = "<span id=\"error-msg\" style=\"color:red;\">The name \"{{ .Name }}\" is already taken. Use a different one.</span>"
		const failed = "<span id=\"error-msg\" style=\"color:green;\">Good, You have successfully registered.</span>"

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
		academicYearStr := r.Form.Get("academic_year")
		academicYear, err := strconv.Atoi(academicYearStr)
		if err != nil {
			fmt.Println("Error converting academic year to integer:", err)
			panic(err)
		} else {
			hitdb.AddStudent(db, r.Form.Get("first_name"), r.Form.Get("last_name"), r.Form.Get("email"), academicYear, r.Form.Get("date_of_birth"))
		}
	})
}
