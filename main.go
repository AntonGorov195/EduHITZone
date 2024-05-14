package main

import (
	hitdb "EduHITZone/src/MySQL"
	spa "EduHITZone/src/page"
	"net/http"
)

// Working on a single page app.
func main() {
	hitdb.ConnectDB()
	http.Handle("/static/", http.FileServer(http.Dir("SPAPublic/")))
	spa.AddPageHandles()

	http.ListenAndServe(":42069", nil)
}
