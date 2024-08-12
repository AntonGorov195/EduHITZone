package spa

import (
	"database/sql"
	"net/http"
)

type Message struct {
	Message string
}

var Messages []Message

func addChatPage(templates *templates, db *sql.DB) {
	Messages = make([]Message, 0)

	http.HandleFunc("POST /chat", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		Messages = append(Messages, Message{r.Form.Get("message")})
		err = templates.Render(w, "chat", Messages)
		if err != nil {
			panic(err)
		}
	})
}
