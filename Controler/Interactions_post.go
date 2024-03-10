package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
	if t.USER.Username != "" {
		if r.Method == "POST" {
			likeValue := 0
			uuid := ""
			if r.FormValue("like") != "" {
				likeValue = 1
				uuid = r.FormValue("like")
			} else if r.FormValue("dislike") != "" {
				likeValue = -1
				uuid = r.FormValue("dislike")
			}

			var previousLike int
			err := db.QueryRow("SELECT likeOrDislike FROM likesFromUser WHERE uuidUser = ? AND uuidLiked = ?", t.USER.Uuid, uuid).Scan(&previousLike)
			switch {
			case err == sql.ErrNoRows:
				_, err := db.Exec("INSERT INTO likesFromUser (uuidUser, uuidLiked, likeOrDislike) VALUES (?, ?, ?)", t.USER.Uuid, uuid, likeValue)
				if err != nil {
					fmt.Println("Erreur lors de l'insertion du like :", err)
					return
				}
			case err != nil:
				fmt.Println("Erreur lors de la recherche du like existant :", err)
				return
			default:
				if previousLike == likeValue {
					_, err := db.Exec("DELETE FROM likesFromUser WHERE uuidUser = ? AND uuidLiked = ?", t.USER.Uuid, uuid)
					if err != nil {
						fmt.Println("Erreur lors de la suppression du like :", err)
						return
					}
				} else {
					_, err := db.Exec("UPDATE likesFromUser SET likeOrDislike = ? WHERE uuidUser = ? AND uuidLiked = ?", likeValue, t.USER.Uuid, uuid)
					if err != nil {
						fmt.Println("Erreur lors de la mise à jour du like :", err)
						return
					}
				}
			}

			var totalLikes int
			err = db.QueryRow("SELECT SUM(likeOrDislike) FROM likesFromUser WHERE uuidLiked = ?", uuid).Scan(&totalLikes)
			if err != nil {
				fmt.Println("Erreur lors du calcul du total des likes :", err)
				return
			}

			_, err = db.Exec("UPDATE topics SET likes = ? WHERE uuid = ?", totalLikes, uuid)
			if err != nil {
				fmt.Println("Erreur lors de la mise à jour du nombre de likes sur le sujet :", err)
				return
			}
		}
	} else {
		fmt.Println("Vous devez être connecté pour aimer ou ne pas aimer un message")
	}
}
