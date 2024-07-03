package spa

import (
	hitdb "EduHITZone/src/MySQL"
	"bytes"
	"database/sql"
	"html/template"
	"net/http"
)

func addAdminHandle(db *sql.DB) {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		courses := hitdb.GetCourses(db)

		var buf bytes.Buffer
		tmpl, err := template.ParseFiles("public/static/views/admin.html", "public/static/views/course-list.html")
		if err != nil {
			panic(err)
		}
		if err := tmpl.Execute(&buf, courses); err != nil {
			panic(err)
		}
		sendViewBuf(w, r, buf)
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
