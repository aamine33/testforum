package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
)

func MessagesSendByUser(db *sql.DB) {
	USER.MessagesSend = nil

	query := fmt.Sprintf("SELECT m.message, m.uuidPath, t.name FROM messages m JOIN topics t ON m.uuidPath = t.uuid WHERE m.owner = '%s'", USER.Username)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		var uuidPath string
		var name string

		err := rows.Scan(&message, &uuidPath, &name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		USER.MessagesSend = append(USER.MessagesSend, t.MessageSend{
			MessageSendByUser: message,
			TopicSentInName:   name,
		})
	}
}
