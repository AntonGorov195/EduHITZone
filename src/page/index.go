package spa

import (
	hitdb "EduHITZone/src/DB"
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type fullPageData struct {
	Header interface{}
	Main   interface{}
	Footer interface{}
}

func (t *templates) RenderFullPage(w io.Writer, page, header, footer, main string, data fullPageData) error {
	var header_buf bytes.Buffer

	if err := t.templates.ExecuteTemplate(&header_buf, header, data.Header); err != nil {
		return err
	}
	var footer_buf bytes.Buffer
	if err := t.templates.ExecuteTemplate(&footer_buf, footer, data.Footer); err != nil {
		return err
	}
	var main_buf bytes.Buffer
	if err := t.templates.ExecuteTemplate(&main_buf, main, data.Main); err != nil {
		return err
	}

	return t.templates.ExecuteTemplate(w, page, struct {
		Header, Footer, Main template.HTML
	}{
		Header: template.HTML(header_buf.String()), Footer: template.HTML(footer_buf.String()), Main: template.HTML(main_buf.String()),
	})
}

func NewTemplates() *templates {
	return &templates{
		templates: template.Must(template.ParseGlob("public/static/views/*.html")),
	}
}

func (t *templates) ConditionalRenderDefault(w io.Writer, is_full_render bool, name string, data interface{}) error {
	if is_full_render {
		return t.RenderFullPage(w, "index", "header", "footer", name, fullPageData{Main: data})
	}
	return t.Render(w, name, data)
}

type adminData struct {
	FormData formData
	Courses  []courseItem
}

func adminDataNew(db *sql.DB) (adminData, error) {
	data := adminData{}
	courses, err := hitdb.GetCourses(db)
	if err != nil {
		return data, err
	}
	data.Courses = toCourseItem(courses, "admin/course")
	data.FormData = formDataNew()
	return data, nil
}

type formData struct {
	Values map[string]string
	Errors map[string]string
}

func formDataNew() formData {
	return formData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func AddPageHandles(db *sql.DB) {
	templates := NewTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			loadEntry(w, r, templates, db)
			return
		}
		loadError(w, r, templates, db)
	})

	// Login
	http.HandleFunc("/login-page", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "login", nil)
		if err != nil {
			panic(err)
		}
	})
	// Submitting login form.
	// TODO: should the login just send you to the courses page or post this request?
	http.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		type LoginData struct {
			Values map[string]string
			Errors map[string]string
		}
		r.ParseForm()
		name := r.Form.Get("name")
		// TODO: Implement passwords...
		password := r.Form.Get("password")
		if name == "admin" {
			w.Header().Add("HX-Push-Url", "/admin")
			w.Header().Add("HX-Location", `{"path":"/admin", "target":"#main-container","swap":"innerHTML"}`)
			data, err := adminDataNew(db)
			if err != nil {
				panic(err)
			}
			err = templates.Render(w, "admin", data)
			if err != nil {
				panic(err)
			}
			return
		}
		students, err := hitdb.GetStudents(db)
		if err != nil {
			w.WriteHeader(502)
			panic(err)
		}

		is_exist := slices.ContainsFunc(students, func(student hitdb.Student) bool {
			return student.FirstName == name
		})

		if !is_exist {
			w.Header().Add("HX-Push-Url", "false")
			w.WriteHeader(422)
			form_data := LoginData{
				Values: make(map[string]string),
				Errors: make(map[string]string),
			}
			form_data.Values["Name"] = name
			form_data.Values["Password"] = password
			form_data.Errors["UserNotFound"] = "Failed finding the account. Make sure you've entered the correct data."
			templates.Render(w, "login", form_data)
			return
		}
		loadSearchCourses(w, r, templates, db)
	})

	// Course search
	http.HandleFunc("/search-course", func(w http.ResponseWriter, r *http.Request) {
		loadSearchCourses(w, r, templates, db)
	})
	// Course
	http.HandleFunc("GET /course", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "course", nil)
		if err != nil {
			panic(err)
		}
	})

	// New account
	http.HandleFunc("GET /new-acc", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "new-acc", nil)
		if err != nil {
			panic(err)
		}
	})
	// TODO: Handle more input field
	http.HandleFunc("POST /new-acc", func(w http.ResponseWriter, r *http.Request) {
		student := hitdb.Student{}

		r.ParseForm()
		academic_year, err := strconv.Atoi(r.Form.Get("academic-year"))
		if err != nil {
			// panic(err)
			student.AcademicYear = 0
			fmt.Println("No academic year")
		} else {
			student.AcademicYear = academic_year
		}
		student.FirstName = r.Form.Get("first-name")
		student.LastName = r.Form.Get("last-name")
		// TODO: Handle birth dates
		student.DateOfBirth = sql.NullString{String: r.Form.Get("birthdate"), Valid: true}
		student, err = hitdb.RegisterStudent(db, student)
		if err != nil {
			// TODO: Handle student failed registration.
			w.WriteHeader(422)
			templates.Render(w, "new-acc-failed", "Failed registering a new student")
			return
		}
		type SuccessData struct {
			FullName string
		}
		templates.Render(w, "new-acc-success", SuccessData{FullName: student.FirstName + " " + student.LastName})
		hitdb.RegisterStudent(db, student)
	})

	// Admin
	http.HandleFunc("GET /admin", func(w http.ResponseWriter, r *http.Request) {
		data, err := adminDataNew(db)
		if err != nil {
			panic(err)
		}
		err = templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "admin", data)
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("POST /admin/course", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		r.URL.Query().Get("name")
		name := r.Form.Get("course-name")
		thumbnail := r.Form.Get("thumbnail")
		course := hitdb.Course{}
		course.Name = name
		course.Thumbnail = thumbnail
		course, err := hitdb.RegisterCourse(db, course)
		if err != nil {
			data, err := adminDataNew(db)
			if err != nil {
				panic(err)
			}
			data.FormData.Values["CourseName"] = name
			data.FormData.Errors["FailedRegistering"] = "Failed to add a new course"
			err = templates.Render(w, "admin", data)
			if err != nil {
				panic(err)
			}
			return
		}

		data, err := adminDataNew(db)
		if err != nil {
			panic(err)
		}
		data.FormData.Values["CourseName"] = name
		err = templates.Render(w, "admin", data)
		if err != nil {
			panic(err)
		}
	})
}

type courseItem struct {
	Name      string
	Thumbnail string
	Href      string
}

func toCourseItem(courses []hitdb.Course, href string) []courseItem {
	items := make([]courseItem, 0, len(courses))
	for _, course := range courses {
		item := courseItem{}
		item.Href = href
		item.Name = course.Name
		item.Thumbnail = course.Thumbnail
		items = append(items, item)
	}
	return items
}

// TODO: Make local anon func in the add pages handles???
func loadSearchCourses(w http.ResponseWriter, r *http.Request, templates *templates, db *sql.DB) {
	courses, err := hitdb.GetCourses(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(courses)
	err = templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "course-list", toCourseItem(courses, "course"))
	if err != nil {
		panic(err)
	}
}
func loadEntry(w http.ResponseWriter, r *http.Request, templates *templates, db *sql.DB) {
	err := templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "error", nil)
	if err != nil {
		panic(err)
	}
}
func loadError(w http.ResponseWriter, r *http.Request, templates *templates, db *sql.DB) {
	err := templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "entry", nil)
	if err != nil {
		panic(err)
	}
}
