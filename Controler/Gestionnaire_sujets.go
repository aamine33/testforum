package forum

import (
	"database/sql"
	"forum/delete"
	"forum/listTopics"
	"forum/login"
	"forum/users"
	"html/template"
	"net/http"
	"strings"
)

func Handler_topics(w http.ResponseWriter, r *http.Request) {
	tmplTopic := template.Must(template.ParseFiles("../static/html/topics.html"))
	tmpl404 := template.Must(template.ParseFiles("../static/html/404.html"))

	users.GetCookieHandler(w, r)

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	if r.FormValue("topic_name") != "" {
		listTopics.AddTopic(r, databaseForum)
	}

	if r.FormValue("input_mail") != "" {
		users.EMAILSTORAGE.Email = r.FormValue("input_mail")
	}

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		login.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		users.Register(r, databaseForum)
	}

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		listTopics.LikesDislikes(r, databaseForum)
	}

	if r.FormValue("delete") != "" {
		delete.DeleteTopic(r, databaseForum)
	}

	if strings.HasPrefix(r.URL.Path, "/topics/category=") {
		listTopics.DisplayTopic(r, databaseForum)
		listTopics.TOPICSANDSESSION.SessionUser = users.USER.Username
		category := strings.TrimPrefix(r.URL.Path, "/topics/category=")
		listTopics.TOPICSANDSESSION.Category = category
		users.GetCookieHandler(w, r)
		tmplTopic.Execute(w, listTopics.TOPICSANDSESSION)
	} else {
		tmpl404.Execute(w, nil)
	}
}
