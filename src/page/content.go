package spa

import (
	hitdb "EduHITZone/src/MySQL"
	"bytes"
	"database/sql"
	"html/template"
	"net/http"
)

func addContentHandle(db *sql.DB) {
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		courses := hitdb.GetCourses(db)

		var buf bytes.Buffer
		tmpl, err := template.ParseFiles("public/static/views/content.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := tmpl.ExecuteTemplate(&buf, "primary", courses); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		sendViewBuf(w, r, buf)
	})

}
