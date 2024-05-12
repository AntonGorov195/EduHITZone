package main

import (
	HITDB "EduHITZone/src/MySQL"
	spa "EduHITZone/src/singlePage"
	"fmt"
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
	spa.AddIndexHandle()
	spa.AddLoginHandle()
	spa.AddContentHandle()

	db := HITDB.ConnectDB()
	HITDB.AddCourse(db, "name")
	fmt.Println(HITDB.GetCourses(db))

	http.ListenAndServe(":42069", nil)
}
