package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

// DeleteTopic supprime un sujet et tous ses messages associés
func SupprimerSujet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Vérifier si l'identifiant du sujet à supprimer est fourni dans la requête
	sujetUUID := r.FormValue("delete")
	if sujetUUID == "" {
		return
	}

	var proprietaire string

	// Récupérer le propriétaire du sujet
	query := "SELECT proprietaire FROM sujets WHERE uuid = ?"
	err := db.QueryRow(query, sujetUUID).Scan(&proprietaire)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du propriétaire du sujet :", err)
		return
	}

	// Vérifier si l'utilisateur est le propriétaire du sujet ou un administrateur
	if t.USER.Username == proprietaire || t.USER.Admin == 1 {
		// Supprimer le sujet
		query = "DELETE FROM sujets WHERE uuid = ?"
		_, err = db.Exec(query, sujetUUID)
		if err != nil {
			fmt.Println("Erreur lors de la suppression du sujet :", err)
			return
		}

		// Supprimer tous les messages associés au sujet
		query = "DELETE FROM messages WHERE uuidChemin = ?"
		_, err = db.Exec(query, sujetUUID)
		if err != nil {
			fmt.Println("Erreur lors de la suppression des messages associés au sujet :", err)
			return
		}

		// Supprimer tous les likes associés au sujet
		query = "DELETE FROM likesUtilisateur WHERE uuidSujet = ?"
		_, err = db.Exec(query, sujetUUID)
		if err != nil {
			fmt.Println("Erreur lors de la suppression des likes associés au sujet :", err)
			return
		}

		fmt.Println("Sujet supprimé avec succès")
	} else {
		fmt.Println("Vous devez être le propriétaire du sujet ou un administrateur pour le supprimer")
	}
}
