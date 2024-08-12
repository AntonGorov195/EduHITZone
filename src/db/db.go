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
	Id       int64 // Id number of the student.
	Username string
	Password []byte
}
type Course struct {
	Id             int64
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

		err = rows.Scan(&course.Id, &course.Name, &course.Thumbnail, &course.VideoLink, &course.Summery, &course.RawTranslation)
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
	insert, err := db.Exec("INSERT INTO courses (name, thumbnail, vid_link, summery, raw_translation) VALUES (?, ?, ?, ?, ?);",
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
	course.Id = id
	return course, nil
}

// TODO:
//
//	func GetCourseById(db *sql.DB, id int) bool {
//		row, err := db.Query("SELECT name FROM courses WHERE students.id = ?;", id)
//		if err != nil {
//			// This shouldn't ever happend.
//			panic(err)
//		}
//		if row.Next() {
//			return true
//		}
//		return false
//	}

func UpdateCourse(db *sql.DB, course Course) error {
	_, err := db.Exec("UPDATE courses SET name = ?, thumbnail = ?, vid_link = ?, summery = ?, raw_translation = ?  WHERE id = ?;",
		course.Name, course.Thumbnail, course.VideoLink, course.Summery, course.RawTranslation, course.Id)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCourse(db *sql.DB, course Course) error {
	_, err := db.Exec("DELETE FROM courses WHERE id = ?;", course.Id)
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

		err = rows.Scan(&student.Id, &student.Username, &student.Password)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

// Register a student into a db. Returns the registered version of the strudent of an error.
// TODO: on error, return empty student or return the unregistered student?
func RegisterStudent(db *sql.DB, student Student) error {
	if student.Id == 0 {
		return errors.New("failed Registering user. Id is invalid (id == 0)")
	}
	_, err := db.Exec("INSERT INTO students (id, username, password) VALUES (?, ?, ?);", student.Id, student.Username, student.Password)
	if err != nil {
		return err
	}
	return nil
}
func GetStudentById(db *sql.DB, id int) (Student, error) {
	var student Student
	row, err := db.Query("SELECT * FROM students WHERE students.id = ?;", id)
	if err != nil {
		return student, err
	}
	if row.Next() {
		err = row.Scan(&student.Id, &student.Username, &student.Password)
		return student, err
	}
	return student, errors.New("user not found")
}

// Returns error if student not found.
func GetStudentByPasswordAndUsername(db *sql.DB, username string, password string) (Student, error) {
	var student Student
	row, err := db.Query("SELECT * FROM hit.students WHERE students.username = ? AND students.password = ?;", username, password)
	if err != nil {
		return student, err
	}
	if row.Next() {
		err = row.Scan(&student.Id, &student.Username, &student.Password)
		return student, err
	}
	return student, errors.New("user not found")
}
func UpdateStudent(db *sql.DB, student Student) error {
	_, err := db.Exec("UPDATE students SET username = ?, password = ? WHERE id = ?;",
		student.Username,
		student.Password,
		student.Id,
	)
	return err
}
func DeleteStudent(db *sql.DB, student Student) error {
	_, err := db.Exec("DELETE FROM students WHERE id = ?;", student.Id)
	return err
}

/*

*/
