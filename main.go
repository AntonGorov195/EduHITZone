package main

import (
	"net/http"

	hitdb "EduHITZone/src/MySQL"
	spa "EduHITZone/src/page"

	_ "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	_ "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
)

// Working on a single page app.
func main() {
	db := hitdb.ConnectDB()
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
