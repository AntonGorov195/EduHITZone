package spa

import (
	hitdb "EduHITZone/src/MySQL"
	"database/sql"
	"html/template"
	"net/http"
)

func addAdminHandle(db *sql.DB) {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "admin")
	})
	http.HandleFunc("/api/v1/admin/form", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		const success = "<span id=\"error-msg\" style=\"color:green;\">Added {{ . }}!</span>"

		name := r.Form.Get("course-name")
		hitdb.AddCourse(db, name)
		tmpl, err := template.New("tmppppp").Parse(success)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, name)
		if err != nil {
			panic(err)
		}
	})
}
