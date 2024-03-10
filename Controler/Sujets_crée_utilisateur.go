package forum

import (
	"database/sql"
	"fmt"
)

func TopicCreatedByUser(db *sql.DB) {
	USER.TopicsCreated = nil
	query := fmt.Sprintf("SELECT name FROM topics WHERE owner = ?", USER.Username)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println(err)
			continue
		}
		USER.TopicsCreated = append(USER.TopicsCreated, name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}
