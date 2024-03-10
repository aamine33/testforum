package forum

import (
	"database/sql"
	"fmt"
	t "forum/listTopics"
	"html/template"
	"net/http"
)

func Handler_EditTopic(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/html/editTopic.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page de modification de sujet", http.StatusInternalServerError)
		return
	}

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	err = t.EditTopic(r, databaseForum)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la modification du sujet : %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page de modification de sujet", http.StatusInternalServerError)
	}
}
