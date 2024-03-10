package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
)

func EditMessage(r *http.Request, db *sql.DB) {
	if r.FormValue("edit") != "" {
		newMessage := r.FormValue("newMessage")

		if len(newMessage) < 10 {
			fmt.Println("Le nouveau message doit contenir au moins 10 caractères.")
			return
		}

		uuid := strings.Split(r.URL.Path, "/")[2]

		var owner string
		query := fmt.Sprintf("SELECT owner FROM messages WHERE uuid = '%s'", uuid)
		row, err := db.Query(query)
		if err != nil {
			fmt.Println("Erreur lors de la récupération du propriétaire du message:", err)
			return
		}
		defer row.Close()
		for row.Next() {
			err := row.Scan(&owner)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du propriétaire du message:", err)
				return
			}
		}

		if t.USER.Username == owner || t.USER.Admin == 1 {
			query := fmt.Sprintf("UPDATE messages SET message = '%s', edited = 1 WHERE uuid = '%s'", newMessage, uuid)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Println("Erreur lors de la mise à jour du message:", err)
				return
			}
			fmt.Println("Le message a été modifié avec succès.")
		} else {
			fmt.Println("Vous devez être le propriétaire du message ou un administrateur pour le modifier.")
		}
	}
}
