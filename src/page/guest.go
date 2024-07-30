package spa

import (
	"database/sql"
	"net/http"
)

func addGuestHandle(templates *templates, db *sql.DB) {
	http.HandleFunc("/guest", func(w http.ResponseWriter, r *http.Request) {

	})
}
