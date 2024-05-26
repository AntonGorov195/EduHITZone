package spa

import (
	"database/sql"
	"net/http"
)

func addIndexHandle(db *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			DrawView(w, r, "error")
			return
		}

		DrawView(w, r, "entry")
	})
}

func AddPageHandles(db *sql.DB) {
	addIndexHandle(db)
	addLoginHandle(db)
	addContentHandle(db)
	addNewAccountHandle(db)
}
