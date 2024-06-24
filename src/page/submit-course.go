package spa

import (
	hitdb "EduHITZone/src/MySQL"
	"database/sql"
	"html/template"
	"net/http"
)

func addCourseSubmitionHandle(db *sql.DB) {
	http.HandleFunc("/submit-course", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "submit-course")
	})
	http.HandleFunc("/api/v1/submit-course/form", func(w http.ResponseWriter, r *http.Request) {
		type ErrorMessageData struct {
			Name string
		}
		r.ParseForm()
		const msg = "<span id=\"error-msg\" style=\"color:white;\">Hi, DB not ready to add {{ .Name }}! This will be fixed soon.</span>"

		name := r.Form.Get("name")
		tmpl, err := template.New("msg-tmpl").Parse(msg)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, ErrorMessageData{Name: name})
		if err != nil {
			panic(err)
		}
		hitdb.AddCourse(db, r.Form.Get("name"))

	})
}
