package main

import (
	spa "EduHITZone/src/page"
	"net/http"
)

//
// func main() {
// 	pages.AddIndexHandles()
// 	fs := http.FileServer(http.Dir("public/"))
// 	http.Handle("/", fs)
//
// 	pages.AddLoginHandles()
// 	http.ListenAndServe(":42069", nil)
// }

// Working on a single page app.
func main() {
	http.Handle("/static/", http.FileServer(http.Dir("SPAPublic/")))
	spa.AddPageHandles()

	http.ListenAndServe(":42069", nil)
}
