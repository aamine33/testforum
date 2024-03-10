package forum

import (
	"database/sql"
	"fmt"
	"forum/delete"
	"forum/login"
	"forum/messages"
	"forum/report"
	"forum/users"
	"html/template"
	"net/http"
	"strings"
)

func Handler_Messages(w http.ResponseWriter, r *http.Request) {
	tmplMessages := template.Must(template.ParseFiles("../static/html/messages.html"))
	tmpl404 := template.Must(template.ParseFiles("../static/html/404.html"))

	users.GetCookieHandler(w, r)

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	if r.FormValue("input_mail") != "" {
		users.EMAILSTORAGE.Email = r.FormValue("input_mail")
	}

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		login.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		users.Register(r, databaseForum)
	}

	if r.FormValue("delete") != "" {
		delete.DeleteMessage(r, databaseForum)
	}

	if r.FormValue("input_newMessage") != "" {
		messages.NewMessage(databaseForum, r)
	}

	if r.FormValue("report") != "" {
		report.ReportMessage(r, databaseForum)
	}

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		messages.LikesDislikes(r, databaseForum)
	}

	if r.FormValue("edit") != "" {
		messages.EditMessage(r, databaseForum)
	}

	messages.MessagesPageDisplay(databaseForum, r)

	var exists bool
	urlName := strings.Split(r.URL.Path, "/")
	if len(urlName) >= 3 {
		newUrlName := strings.TrimSpace(urlName[2])
		query := "SELECT name FROM topics WHERE name = ?"
		row, err := databaseForum.Query(query, newUrlName)
		if err != nil {
			fmt.Println(err)
		} else {
			defer row.Close()
			for row.Next() {
				var name string
				err = row.Scan(&name)
				if err != nil {
					fmt.Println(err)
				} else {
					if name == newUrlName {
						exists = true
					}
				}
			}
		}
	}

	if exists {
		messages.MESSAGES.SessionUser = users.USER.Username
		users.GetCookieHandler(w, r)
		tmplMessages.Execute(w, messages.MESSAGES)
	} else {
		tmpl404.Execute(w, nil)
	}
}
