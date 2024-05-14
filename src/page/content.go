package spa

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func AddContentHandle() {
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		tmpl, err := template.ParseFiles("SPAPublic/static/views/content.html")
		if err != nil {
			fmt.Println("Failed to parse content.html template")
			panic(err.Error())
		}
		if err := tmpl.Execute(&buf, nil); err != nil {
			fmt.Println("Failed to execute content.html template")
			panic(err.Error())
		}
		Render(w, r, buf)
	})
}
