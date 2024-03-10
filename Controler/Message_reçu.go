package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
)

func MpSendOrReceivedByUser(db *sql.DB) {
	USER.PrivateMessages = nil

	query := fmt.Sprintf("SELECT message, CASE WHEN user1 = ? THEN user2 ELSE user1 END AS other_user FROM mp WHERE user1 = ? OR user2 = ?")
	rows, err := db.Query(query, USER.Username, USER.Username, USER.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		var otherUser string

		err := rows.Scan(&message, &otherUser)
		if err != nil {
			fmt.Println(err)
			continue
		}

		USER.PrivateMessages = append(USER.PrivateMessages, t.PrivateMessage{
			PrivateMessage:        message,
			PrivateMessage2ndUser: otherUser,
		})
	}
}
