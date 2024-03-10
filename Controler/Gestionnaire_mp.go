package forum

import (
	"database/sql"
	"forum/mp"
	"forum/users"
	"html/template"
	"net/http"
)

func Handler_Mp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/mp.html"))

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	users.GetCookieHandler(w, r)

	if r.FormValue("mpMessage") != "" {
		mp.AddMp(r, databaseForum)
	}

	mp.DisplayMp(r, databaseForum)

	tmpl.Execute(w, mp.MPSANDTOWHO)
}
