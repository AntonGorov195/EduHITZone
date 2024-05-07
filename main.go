package main

import (
	loginPage "EduHITZone/src/page"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	loginPage.AddHandles()
	http.ListenAndServe(":42069", nil)
}
