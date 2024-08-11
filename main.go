package main

import (
	"net/http"

	hitdb "EduHITZone/src/DB"
	spa "EduHITZone/src/page"
)

// Working on a single page app.
func main() {
	db, err := hitdb.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/static/"))))
	spa.AddPageHandles(db)

	http.ListenAndServe(":42069", nil)

}

/*hitdb.ConnectDB()
http.Handle("/static/", http.FileServer(http.Dir("SPAPublic/")))
spa.AddPageHandles()
fmt.Println("Hello")
insert, err := db.query("INSERT INTO users VALUES('OMER')")
if err != nil {
	panic(err.Error())
}
defer insert.Close()*/
