package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Course struct {
	Id        int
	Name      string
	Thumbnail string
	Href      string
}
type CourseList struct {
	Name    string
	Courses []Course
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "wO2F56oXL4Wy4D8RJFlP"
	dbname   = "edu_hit_zone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		course := Course{Name: r.Form.Get("name"), Thumbnail: r.Form.Get("thumbnail"), Href: r.Form.Get("thumbnail")}

		_, err = db.Exec("INSERT INTO course(name,thumbnail,page) VALUES($1,$2,$3)", course.Name, course.Thumbnail, course.Href)
		if err != nil {
			panic(err)
		}
		tmpl := template.Must(template.ParseFiles("public/contentPage/content.html"))
		if err := tmpl.ExecuteTemplate(w, "course_preview", course); err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/contentPage/content.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("public/contentPage/content.html"))

		courses := []Course{}
		rows, err := db.Query("SELECT * FROM course")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var course Course
			if err := rows.Scan(&course.Id, &course.Name, &course.Thumbnail, &course.Href); err != nil {
				panic(err)
			}
			courses = append(courses, course)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		index := map[string][]CourseList{
			"CoursesLists": {
				{
					Name: "Popular", Courses: courses,
				},
			},
		}
		if err := tmpl.ExecuteTemplate(w, "index", index); err != nil {
			panic(err)
		}
	})
	fmt.Println("Start Listen")
	http.ListenAndServe(":8080", nil)
}
