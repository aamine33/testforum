package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func ReporterMessage(r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		var uuidReported string
		var dejaSignale bool

		uuid := r.FormValue("report")

		if t.USER.Username != "" {
			query := fmt.Sprintf("SELECT uuidReported FROM reports WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := db.Query(query)
			if err != nil {
				fmt.Println("Erreur lors de la requête à la base de données :", err)
			} else {
				defer row.Close()
				for row.Next() {
					err = row.Scan(&uuidReported)
					if err != nil {
						fmt.Println("Erreur lors de la lecture de la ligne :", err)
						continue
					}
					if uuidReported == uuid {
						dejaSignale = true
						break
					}
				}
			}

			if !dejaSignale {
				query = fmt.Sprintf("INSERT INTO reports(uuidUser, uuidReported) VALUES ('%s', '%s')", t.USER.Uuid, uuid)
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println("Erreur lors de l'insertion du signalement dans la base de données :", err)
				}

				query = fmt.Sprintf("UPDATE messages SET report = report + 1 WHERE uuid = '%s'", uuid)
				_, err = db.Exec(query)
				if err != nil {
					fmt.Println("Erreur lors de la mise à jour du nombre de signalements du message :", err)
				}
			} else {
				fmt.Println("Message déjà signalé")
			}
		} else {
			fmt.Println("Vous devez être connecté pour signaler un message")
		}

		_, err := db.Exec("DELETE FROM messages WHERE report >= 5")
		if err != nil {
			fmt.Println("Erreur lors de la suppression des messages signalés :", err)
		}
	}
}
