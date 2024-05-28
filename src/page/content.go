package spa

import (
	"database/sql"
	"net/http"
)

func addContentHandle(db *sql.DB) {
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		drawView(w, r, "content")
	})
}
