package spa

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func AddIndexHandle() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		if r.URL.Path != "/" {
			tmpl, err := template.ParseFiles("SPAPublic/error.html")
			if err != nil {
				fmt.Println("Failed to parse error.html template")
				panic(err.Error())
			}
			if err := tmpl.Execute(&buf, nil); err != nil {
				fmt.Println("Failed to execute error.html template")
				panic(err.Error())
			}
			Render(w, r, buf)
			return
		}

		tmpl, err := template.ParseFiles("SPAPublic/static/views/entry.html")
		if err != nil {
			fmt.Println("Failed to parse entry.html template")
			panic(err.Error())
		}
		if err := tmpl.Execute(&buf, nil); err != nil {
			fmt.Println("Failed to execute entry.html template")
			panic(err.Error())
		}
		Render(w, r, buf)
	})
}
