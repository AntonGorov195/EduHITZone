package spa

import (
	hitdb "EduHITZone/src/DB"
	ai "EduHITZone/src/ai"
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
		// TODO: Implement passwords...
		password := r.Form.Get("password")
		username := r.Form.Get("username")
		if username == "admin" {
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
		_, err := hitdb.GetStudentByPasswordAndUsername(db, username, password)
		if err != nil {
			// TODO: better error handling
			w.Header().Add("HX-Push-Url", "false")
			w.WriteHeader(422)
			form_data := LoginData{
				Values: make(map[string]string),
				Errors: make(map[string]string),
			}
			form_data.Values["Name"] = username
			form_data.Values["Password"] = password
			form_data.Errors["UserNotFound"] = "סיסמה או שם משתמש לא נכון"
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
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			panic(err)
		}
		courses, err := hitdb.GetCourses(db)
		if err != nil {
			panic(err)
		}
		course := courses[slices.IndexFunc(courses, func(course hitdb.Course) bool {
			return course.Id == int64(id)
		})]

		err = templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "course", struct {
			CourseId int
			Summary  string
			Video    string
		}{CourseId: id, Summary: course.Summery.String, Video: course.VideoLink.String})

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
	http.HandleFunc("POST /new-acc", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		password := []byte(r.Form.Get("password"))
		username := r.Form.Get("username")
		id, err := strconv.Atoi(r.Form.Get("user-id"))
		if err != nil {
			w.Header().Add("HX-Push-Url", "false")
			w.Header().Add("HX-Location", `{"path":"/admin", "target":"#error-msg","swap":"innerHTML"}`)
			panic(err)
		}
		student := hitdb.Student{Id: int64(id), Username: username, Password: password}
		err = hitdb.RegisterStudent(db, student)
		if err != nil {
			// TODO: Handle student failed registration.
			w.WriteHeader(422)
			templates.Render(w, "new-acc-failed", "Failed registering a new student")
			panic(err)
		}
		loadSearchCourses(w, r, templates, db)
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
		course := hitdb.Course{}
		err := r.ParseMultipartForm(99999999999999999)
		if err != nil {
			panic(err)
		}
		name := r.Form.Get("course-name")
		course.Name = name
		fmt.Println(name)
		{
			file, _, err := r.FormFile("thumbnail")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			// Create the uploads directory if it doesn't exist
			uploadDir := "public/static/thumbnails/"
			if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
				os.Mkdir(uploadDir, 0755)
			}

			// Save the uploaded file to the server
			filePath := filepath.Join(uploadDir, name)
			out, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				panic(err)
			}
			course.Thumbnail = "/static/thumbnails/" + name
		}
		{
			file, _, err := r.FormFile("video")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			// Create the uploads directory if it doesn't exist
			uploadDir := "public/static/videos/"
			if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
				os.Mkdir(uploadDir, 0755)
			}

			// Save the uploaded file to the server
			filePath := filepath.Join(uploadDir, name)
			out, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				panic(err)
			}
			course.VideoLink.String = "/static/videos/" + name
			course.VideoLink.Valid = true
		}
		course, err = hitdb.RegisterCourse(db, course)
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
	http.HandleFunc("GET /admin/course", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			panic(err)
		}
		courses, err := hitdb.GetCourses(db)
		if err != nil {
			panic(err)
		}
		course := courses[slices.IndexFunc(courses, func(course hitdb.Course) bool {
			return course.Id == int64(id)
		})]

		err = templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "course-admin", struct {
			CourseId int
			Summary  string
			Video    string
		}{CourseId: id, Summary: course.Summery.String, Video: course.VideoLink.String})

		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("PUT /admin/course", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		courses, err := hitdb.GetCourses(db)
		if err != nil {
			panic(err)
		}
		course := courses[slices.IndexFunc(courses, func(course hitdb.Course) bool {
			return course.Id == int64(id)
		})]
		course.Summery = sql.NullString{String: r.Form.Get("summary"), Valid: true}
		err = hitdb.UpdateCourse(db, course)

		if err != nil {
			panic(err)
		}
		err = templates.ConditionalRenderDefault(w, r.Header.Get("HX-Request") == "", "course-admin", struct {
			CourseId int
			Summary  string
			Video    string
		}{CourseId: id, Summary: course.Summery.String, Video: course.VideoLink.String})

		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("GET /admin/course/summary", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			panic(err)
		}
		courses, err := hitdb.GetCourses(db)
		if err != nil {
			panic(err)
		}
		course := courses[slices.IndexFunc(courses, func(course hitdb.Course) bool {
			return course.Id == int64(id)
		})]
		summary, _, err := ai.GenerateAIContent("public" + course.VideoLink.String)
		if err != nil {
			panic(err)
		}

		course.Summery = sql.NullString{String: summary, Valid: true}
		hitdb.UpdateCourse(db, course)
		io.WriteString(w, summary)
	})
	addChatPage(templates, db)
}

type courseItem struct {
	Id        int64
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
		item.Id = course.Id
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
