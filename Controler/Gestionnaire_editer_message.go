package forum

import (
	"database/sql"
	"fmt"
	t "forum/messages"
	"html/template"
	"net/http"
)

func HandlerEditMessage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/html/editMessage.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture de la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = t.EditMessage(r, db)
	if err != nil {
		fmt.Println("Erreur lors de la modification du message:", err)
		http.Error(w, "Erreur lors de la modification du message", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Erreur lors de l'affichage de la page de modification de message:", err)
		http.Error(w, "Erreur lors de l'affichage de la page de modification de message", http.StatusInternalServerError)
	}
}
