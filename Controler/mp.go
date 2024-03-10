package forum

import (
	"database/sql"
	t "forum/users"
	"net/http"
	"strings"
	"time"
)

func AddMp(r *http.Request, db *sql.DB) {
	mpMessage := r.FormValue("mpMessage")
	if mpMessage == "" {
		http.Error(w, "Empty message", http.StatusBadRequest)
		return
	}

	user2 := strings.Split(r.URL.Path, "/")
	if t.USER.Username == user2[2] {
		http.Error(w, "Cannot send message to yourself", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO mp(user1, user2, creationDate, message) VALUES (?, ?, ?, ?)"
	creationDate := time.Now()
	_, err := db.Exec(query, t.USER.Username, user2[2], creationDate, mpMessage)
	if err != nil {
		http.Error(w, "Failed to add message", http.StatusInternalServerError)
		return
	}
}
