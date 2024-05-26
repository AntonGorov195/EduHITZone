package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	hitdb "EduHITZone/src/MySQL"
	spa "EduHITZone/src/page"
)

type Student struct {
	student_id    int
	first_name    string
	last_name     string
	email         string
	academic_year int
	date_of_birth string
	phone_number  int
}

// Working on a single page app.
func main() {
	db := hitdb.ConnectDB()
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("SPAPublic/"))))
	spa.AddPageHandles()

	err := AddStudent(db, "yovel")
	//students, err := GetStudents(db)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}
	//fmt.Println("students list:", students)
	http.ListenAndServe(":42069", nil)

}
func AddStudent(db *sql.DB, first_name string) error {
	_, err := db.Exec("INSERT INTO students (first_name) VALUES (?);", first_name)
	//_, err := db.Exec("DELETE FROM students WHERE students_id IS NULL;")
	if err != nil {
		return fmt.Errorf("failed to insert student: %w", err)
	}
	fmt.Println("Successfully added student:", first_name)
	return nil
}
func GetStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM students;")
	if err != nil {
		return nil, fmt.Errorf("failed to get students: %w", err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.student_id, &student.first_name, &student.last_name, &student.email, &student.academic_year, &student.date_of_birth, &student.phone_number)

		if err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return students, nil
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
