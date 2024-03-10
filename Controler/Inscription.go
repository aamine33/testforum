package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/structs"
	t "forum/views"
	"net/http"
	"time"
)

var EMAILSTORAGE t2.EmailStorage

func Register(r *http.Request, database *sql.DB) {
	if r.Method == http.MethodPost {
		fmt.Println("Start registration process")
		fmt.Println("New POST: (register) ")

		username := r.FormValue("input_username")
		password := r.FormValue("input_password")
		password2 := r.FormValue("input_password2")
		birthDay := r.FormValue("input_birthDay")
		mail := r.FormValue("input_mail")
		creationDate := time.Now().String()

		if password != password2 {
			fmt.Println("Passwords don't match")
			return
		}

		if len(username) < 5 || len(username) > 14 {
			fmt.Println("Invalid username length")
			return
		}

		if !t.CheckPassword(password) {
			fmt.Println("Invalid password")
			return
		}

		if !t.CheckMail(mail) {
			fmt.Println("Invalid email format")
			return
		}

		AddUsers(database, username, t.Hash(password), mail, creationDate, birthDay)
	}
}
