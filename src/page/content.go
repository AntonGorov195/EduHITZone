package page

func AddContentHandles() {
	// 	http.HandleFunc("/login/login.html", func(w http.ResponseWriter, r *http.Request) {
	// 		tmpl_main, _ := template.ParseFiles("public/login/login.html")
	// 		tmpl, _ := templazte.ParseFiles("public/template.html")
	//
	// 		var buf bytes.Buffer
	// 		// Reads the main content of the page.
	// 		if err := tmpl_main.Execute(&buf, nil); err != nil {
	// 			panic(err)
	// 		}
	// 		page_main_content := buf.String()
	//
	// 		// Add the main content to the template, which will be rendered.
	// 		tmpl.ExecuteTemplate(w, "primary", PageTemplateData{
	// 			[]template.HTML{
	// 				GetCSSElement(tmpl, "login.css"),
	// 				GetCSSElement(tmpl, "../common.css"),
	// 			},
	// 			template.HTML(page_main_content),
	// 			"Login"})
	// 	})
}
