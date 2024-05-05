package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl, _ := template.ParseFiles("public/index.html")
	// 	tmpl.Execute(w, nil)
	// })

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":42069", nil)
}
