package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/structs"
	t "forum/views"
	"net/http"
)

func UserEdit(r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		if USER.Username != "" {
			username := r.FormValue("username")
			email := r.FormValue("email")

			if username != "" {
				if len(username) < 5 || len(username) > 14 {
					fmt.Println("invalid username")
				} else {
					if err := updateUserUsername(db, username); err != nil {
						fmt.Println("Error updating username:", err)
					}
				}
			}

			if email != "" {
				if !t.CheckMail(email) {
					fmt.Println("Invalid email")
				} else {
					if err := updateUserEmail(db, email); err != nil {
						fmt.Println("Error updating email:", err)
					}
				}
			}

			Logout(r)
		} else {
			fmt.Println("You need to be logged in to edit your account")
		}
	}
}

func updateUserUsername(db *sql.DB, username string) error {
	query := "UPDATE users SET username = ? WHERE uuid = ?"
	if _, err := db.Exec(query, username, USER.Uuid); err != nil {
		return err
	}

	queries := []string{
		"UPDATE messages SET owner = ? WHERE owner = ?",
		"UPDATE topics SET owner = ? WHERE owner = ?",
	}

	for _, q := range queries {
		if _, err := db.Exec(q, username, USER.Username); err != nil {
			return err
		}
	}

	return nil
}

func updateUserEmail(db *sql.DB, email string) error {
	query := "UPDATE users SET email = ? WHERE uuid = ?"
	_, err := db.Exec(query, email, USER.Uuid)
	return err
}
