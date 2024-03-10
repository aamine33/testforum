package forum

import (
	"database/sql"
	"fmt"
	"forum/logOutSessionHtml"
	"forum/users"
	"html/template"
	"net/http"
)

func Handler_profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/profil.html"))

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	users.GetCookieHandler(w, r)

	if r.FormValue("logOutButton") != "" {
		users.Logout(r)
		users.LogOutCookie(r, w)
		logOutSessionHtml.LogOutSession()
		users.USER.Username = ""
	} else if r.FormValue("delete") != "" {
		users.DeleteAccount(r, databaseForum, w)
	} else if r.FormValue("username") != "" || r.FormValue("email") != "" {
		users.UserEdit(r, databaseForum)
	}

	users.MpSendOrReceivedByUser(databaseForum)

	users.MessagesSendByUser(databaseForum)

	users.TopicCreatedByUser(databaseForum)

	fmt.Println(users.USER.Username, "username")
	tmpl.Execute(w, users.USER)
}
