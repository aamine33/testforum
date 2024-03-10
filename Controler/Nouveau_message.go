package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewMessage(db *sql.DB, r *http.Request) {
	uuidMessage := uuid.New()

	topicName := strings.Split(r.URL.Path, "/")

	if r.Method == "POST" {
		var uuidPath string
		query := fmt.Sprintf("SELECT uuid FROM topics WHERE name = '%s'", topicName[2])
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer row.Close()
		for row.Next() {
			err := row.Scan(&uuidPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		message := r.FormValue("input_newMessage")

		if len(message) < 2 {
			fmt.Println("not enough characters to post a message")
			return
		} else if t.USER.Username == "" {
			fmt.Println("you need to be logged in to post a message")
			return
		}

		creationDate := time.Now().String()

		newMessageQuery := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, edited, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
		queryMessage, err := db.Prepare(newMessageQuery)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = queryMessage.Exec(message, creationDate, t.USER.Username, 0, uuidPath, 0, 0, uuidMessage)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			fmt.Println("new message")
			query2 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts + 1, lastPost = '%s' WHERE uuid = '%s'", creationDate, uuidPath)
			_, err := db.Exec(query2)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
