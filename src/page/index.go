package spa

import (
	"database/sql"
	"net/http"
)

func addIndexHandle(db *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			drawView(w, r, "error")
			return
		}

		drawView(w, r, "entry")
	})
}

func AddPageHandles(db *sql.DB) {
	addIndexHandle(db)
	addLoginHandle(db)
	addContentHandle(db)
	addNewAccountHandle(db)
	addCourseSubmitionHandle(db)
}
