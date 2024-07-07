package eduhitdb

import "testing"

// This will be divided into multple unit test later.
func TestCourse(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	courses, err := GetCourses(db)
	if err != nil {
		t.Error(err)
	}
	t.Log("Current courses: ", courses)
	new_course := Course{}
	new_course.Name = "Test Course"
	new_course, err = RegisterCourse(db, new_course)
	if err != nil {
		t.Error(err)
	}
	courses, err = GetCourses(db)
	if err != nil {
		t.Error(err)
	}
	t.Log("After adding courses: ", courses)

	new_course.Name = "New Test Course"
	err = UpdateCourse(db, new_course)
	if err != nil {
		t.Error(err)
	}

	err = DeleteCourse(db, new_course)
	if err != nil {
		t.Error(err)
	}
	t.Log("After deleting courses: ", courses)
}
func TestStudent(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	students, err := GetStudents(db)
	if err != nil {
		t.Error(err)
	}
	t.Log("Current students: ", students)
	new_student := Student{}
	new_student.FirstName = "Test"
	new_student.LastName = "Student"
	new_student, err = RegisterStudent(db, new_student)
	if err != nil {
		t.Error(err)
	}
	students, err = GetStudents(db)
	if err != nil {
		t.Error(err)
	}
	t.Log("After adding a student: ", students)

	new_student.FirstName = "New Test"
	err = UpdateStudent(db, new_student)
	if err != nil {
		t.Error(err)
	}

	err = DeleteStudent(db, new_student)
	if err != nil {
		t.Error(err)
	}
	t.Log("After deleting test student: ", new_student)
}
