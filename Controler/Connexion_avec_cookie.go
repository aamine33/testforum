package forum

import (
	"database/sql"
	"fmt"
	t "forum/views"
)

func LoginWithCookie(uuidUser string) {
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")
	var username string
	var email string
	var creationDate string
	var birthDate string
	var admin int
	query := fmt.Sprintf("SELECT username, email, creationDate, birthDate, admin FROM users WHERE uuid = '%s'", uuidUser)
	row, err := databaseForum.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&username, &email, &creationDate, &birthDate, &admin)
		if err != nil {
			fmt.Println(err)
			return
		}

		USER.Username = username
		USER.Uuid = uuidUser
		USER.Email = email
		USER.CreationDate = t.DisplayTime(creationDate)
		USER.Admin = admin
		USER.BirthDate = birthDate
	}
}
