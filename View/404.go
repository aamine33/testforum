package forum

import (
	"html/template"
	"net/http"
)

func Handler404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/html/404.html")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la page d'erreur 404", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page d'erreur 404", http.StatusInternalServerError)
	}
}
