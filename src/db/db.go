package eduhitdb

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Change this to your mysql password.
var password string = os.Getenv("MYSQL_PASSWORD")

func ConnectDB() (*sql.DB, error) {
	// password := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/hit")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Student struct {
	id           int64
	FirstName    string
	LastName     string
	Email        sql.NullString
	DateOfBirth  sql.NullString
	AcademicYear int
}
type Course struct {
	id             int64
	Name           string
	Thumbnail      string
	VideoLink      sql.NullString
	Summery        sql.NullString
	RawTranslation sql.NullString
}

func GetCourses(db *sql.DB) ([]Course, error) {
	rows, err := db.Query("SELECT * FROM courses;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course

		err = rows.Scan(&course.id, &course.Name, &course.Thumbnail, &course.VideoLink, &course.Summery, &course.RawTranslation)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

// Register a course to a db. Returns the registered version of the course of an error.
// TODO: on error, return empty course or return the unregistered course?
func RegisterCourse(db *sql.DB, course Course) (Course, error) {
	if IsCourseRegistered(db, course) {
		return course, errors.New("attempt to register an existing course. try using UpdateCourse")
	}
	insert, err := db.Exec("INSERT INTO courses (name, thumbnail, vid_link, summery, raw_translation) VALUES (?,?,?,?,?);",
		course.Name,
		course.Thumbnail,
		course.VideoLink,
		course.Summery,
		course.RawTranslation,
	)
	if err != nil {
		return course, err
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return course, err
	}
	course.id = id
	return course, nil
}
func IsCourseRegistered(db *sql.DB, course Course) bool {
	// TODO: Check if the course is already in the db.
	return course.id != 0
}
func UpdateCourse(db *sql.DB, course Course) error {
	_, err := db.Exec("UPDATE courses SET name = ?, thumbnail = ?, vid_link=?, summery = ?, raw_translation = ?  WHERE id = ?;",
		course.Name, course.Thumbnail, course.VideoLink, course.Summery, course.RawTranslation, course.id)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCourse(db *sql.DB, course Course) error {
	_, err := db.Exec("DELETE FROM courses WHERE id = ?;", course.id)
	if err != nil {
		return err
	}
	return nil
}
func GetStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM students;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var student Student

		err = rows.Scan(&student.id, &student.FirstName, &student.LastName, &student.Email, &student.AcademicYear, &student.DateOfBirth)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

// Register a student into a db. Returns the registered version of the strudent of an error.
// TODO: on error, return empty student or return the unregistered student?
func RegisterStudent(db *sql.DB, student Student) (Student, error) {
	if IsStudentRegistered(db, student) {
		return student, errors.New("attempt to register an existing student. try using UpdateStudent")
	}
	insert, err := db.Exec("INSERT INTO students (first_name, last_name, email, academic_year, date_of_birth) VALUES (?, ?, ?, ?, ?);", student.FirstName, student.LastName, student.Email,
		student.AcademicYear, student.DateOfBirth)
	if err != nil {
		return student, err
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return student, err
	}
	student.id = id
	return student, nil
}
func IsStudentRegistered(db *sql.DB, student Student) bool {
	// TODO: Check if the student is already in the db.
	return student.id != 0
}
func UpdateStudent(db *sql.DB, student Student) error {
	_, err := db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, date_of_birth=?, academic_year = ? WHERE id = ?;",
		student.FirstName,
		student.LastName,
		student.Email,
		student.DateOfBirth,
		student.AcademicYear,
		student.id,
	)
	if err != nil {
		return err
	}
	return nil
}
func DeleteStudent(db *sql.DB, student Student) error {
	_, err := db.Exec("DELETE FROM students WHERE id = ?;", student.id)
	if err != nil {
		return err
	}
	return nil
}
