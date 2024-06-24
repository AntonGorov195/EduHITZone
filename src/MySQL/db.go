package hitdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:lRQmFIhGzeEADJOYtm3l@tcp(localhost:3306)/hit")
	if err != nil {
		fmt.Println("Error opening db")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed pinging the db")
		panic(err.Error())
	}
	fmt.Println("Connected to the db!")
	return db
}

type Students struct {
	Student_id    int
	First_name    string
	Last_name     string
	Email         string
	Academic_year int
	Date_of_birth string
}
type Course struct {
	ID        int
	Name      string
	Thumbnail sql.NullString
	Href      sql.NullString
}

func GetCourses(db *sql.DB) []Course {
	rows, err := db.Query("SELECT * FROM courses;")
	if err != nil {
		fmt.Println("Failed getting rows from db")
		panic(err.Error())
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course

		err = rows.Scan(&course.ID, &course.Name, &course.Thumbnail, &course.Href)
		if err != nil {
			fmt.Println("Failed scanning row")
			panic(err.Error())
		}
		courses = append(courses, course)
	}

	return courses
}
func AddCourse(db *sql.DB, name string) {
	insert, err := db.Query("INSERT INTO courses (name) VALUES (?);", name)
	if err != nil {
		fmt.Println("Failed adding course db")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Succesfuly add course")

}
func UpdateCourse(db *sql.DB, id int, name string) {
	update, err := db.Query("UPDATE courses SET course_name = ? WHERE course_id = ?;", name, id)
	if err != nil {
		fmt.Println("Failed updating course in db")
		panic(err.Error())
	}
	defer update.Close()
	fmt.Println("Successfully updated course")
}

func DeleteCourse(db *sql.DB, id int) {
	delete, err := db.Query("DELETE FROM courses WHERE course_id = ?;", id)
	if err != nil {
		fmt.Println("Failed deleting course from db")
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Println("Successfully deleted course")
}
func GetStudents(db *sql.DB) []Students {
	rows, err := db.Query("SELECT * FROM students;")
	if err != nil {
		fmt.Println("Failed getting rows from db")
		panic(err.Error())
	}
	defer rows.Close()
	var students []Students
	for rows.Next() {
		var student Students

		err = rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &student.Academic_year, &student.Date_of_birth)
		if err != nil {
			fmt.Println("Failed scanning row")
			panic(err.Error())
		}
		students = append(students, student)
	}

	return students
}

func AddStudent(db *sql.DB, firstName string, lastName string, email string, academicYear int, dateOfBirth string) {
	insert, err := db.Query("INSERT INTO students (first_name, last_name, email, academic_year, date_of_birth) VALUES (?, ?, ?, ?, ?);", firstName, lastName, email, academicYear, dateOfBirth)
	if err != nil {
		fmt.Println("Failed adding student to db")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully added student")
}
func UpdateStudent(db *sql.DB, id int, firstName string, lastName string, email string, academicYear int, dateOfBirth string) {
	update, err := db.Query("UPDATE students SET first_name = ?, last_name = ?, email = ?, academic_year = ?, date_of_birth = ? WHERE student_id = ?;", firstName, lastName, email, academicYear, dateOfBirth, id)
	if err != nil {
		fmt.Println("Failed updating student in db")
		panic(err.Error())
	}
	defer update.Close()
	fmt.Println("Successfully updated student")
}

func DeleteStudent(db *sql.DB, id int) {
	delete, err := db.Query("DELETE FROM students WHERE student_id = ?;", id)
	if err != nil {
		fmt.Println("Failed deleting student from db")
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Println("Successfully deleted student")
}
