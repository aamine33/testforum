package forum

import (
	"database/sql"
	"forum/delete"
	"forum/home"
	"forum/listTopics"
	"forum/login"
	"forum/users"
	"html/template"
	"net/http"
)

func Handler_Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/html/home.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page d'accueil", http.StatusInternalServerError)
		return
	}

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		listTopics.LikesDislikes(r, databaseForum)
	}

	if r.FormValue("delete") != "" {
		delete.DeleteTopic(r, databaseForum)
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

	home.GetRandomMessages(databaseForum, r)

	err = tmpl.Execute(w, users.TOPICSANDSESSION)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page d'accueil", http.StatusInternalServerError)
	}
}
