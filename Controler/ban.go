package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/profil"
	t "forum/users"
)

func Ban(db *sql.DB) {
	if t.USER.Admin != 1 {
		fmt.Println("You need to be an admin to ban someone.")
		return
	}

	query := "UPDATE users SET ban = 1 WHERE username = ?"
	_, err := db.Exec(query, t2.PUBLICUSER.Username)
	if err != nil {
		fmt.Println("Error banning user:", err)
		return
	}

	fmt.Println("User", t2.PUBLICUSER.Username, "has been banned.")
}
