package spa

import "net/http"

func AddNewAccountHandle() {
	http.HandleFunc("/new-acc", func(w http.ResponseWriter, r *http.Request) {
		DrawView(w, r, "new-acc")
	})
}
