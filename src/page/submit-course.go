package spa

import (
	"database/sql"
	"net/http"
)

func addCourseSubmitionHandle(db *sql.DB) {
	http.HandleFunc("/submit-course", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "submit-course")
	})
}
