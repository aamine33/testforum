package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/profil"
	t "forum/users"
	"net/http"
)

func ReporterUtilisateur(r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		var uuidReported string
		var alreadyReported bool

		uuid := r.FormValue("report")

		if t.USER.Username != "" && t.USER.Username != uuid {
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
						alreadyReported = true
						break
					}
				}
			}

			if !alreadyReported {
				query = fmt.Sprintf("INSERT INTO reports(uuidUser, uuidReported) VALUES ('%s', '%s')", t.USER.Uuid, uuid)
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println("Erreur lors de l'insertion du signalement dans la base de données :", err)
				}

				report := "reports"
				query = fmt.Sprintf("UPDATE users SET %s = %s + 1 WHERE username = '%s'", report, report, uuid)
				_, err = db.Exec(query)
				if err != nil {
					fmt.Println("Erreur lors de la mise à jour du nombre de signalements de l'utilisateur :", err)
				}
			} else {
				fmt.Println("Utilisateur déjà signalé")
			}
		} else {
			fmt.Println("Vous devez être connecté pour signaler un utilisateur ou vous ne pouvez pas vous signaler vous-même")
		}

		if t2.PUBLICUSER.Reports >= 9 {
			query := fmt.Sprintf("UPDATE users SET ban = 1 WHERE username = '%s'", t2.PUBLICUSER.Username)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Println("Erreur lors du bannissement de l'utilisateur :", err)
			}
		}
	}
}
