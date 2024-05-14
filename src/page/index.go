package spa

import (
	"net/http"
)

func AddIndexHandle() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			DrawView(w, r, "error")
			return
		}

		DrawView(w, r, "entry")
	})
}

func AddPageHandles() {
	AddIndexHandle()
	AddLoginHandle()
	AddContentHandle()
	AddNewAccountHandle()
}
