package hitdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:ofekbiton1234@tcp(localhost:3306)/hit")
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

func AddStudents(db *sql.DB)
