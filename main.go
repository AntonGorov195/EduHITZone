package main

import (
	pages "EduHITZone/src/page"
	"net/http"
)

func main() {
	pages.AddIndexHandles()
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)

	pages.AddLoginHandles()
	http.ListenAndServe(":42069", nil)
}
