package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
	if t.USER.Username != "" {
		var newLike int
		var uuid string
		var uuidLiked string
		var alreadyLiked bool
		var previousLike int
		var likeOrDislike int
		var allLikes int
		var query string

		if r.Method == "POST" {
			if r.FormValue("like") != "" {
				newLike = 1
				uuid = r.FormValue("like")
			}
			if r.FormValue("dislike") != "" {
				newLike = -1
				uuid = r.FormValue("dislike")
			}

			query = fmt.Sprintf("SELECT uuidLiked, likeOrDislike FROM likesFromUser WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := db.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					defer row.Close()
					row.Scan(&uuidLiked, &previousLike)
					if uuidLiked == uuid {
						alreadyLiked = true
					}
				}

				if alreadyLiked {
					if previousLike == newLike {
						query = fmt.Sprintf("DELETE FROM likesFromUser WHERE uuidLiked = '%s' AND uuidUser = '%s'", uuidLiked, t.USER.Uuid)
					} else {
						query = fmt.Sprintf("UPDATE likesFromUser SET likeOrDislike = '%d' WHERE uuidLiked = '%s' AND uuidUser = '%s'", newLike, uuid, t.USER.Uuid)
					}
				} else {
					query = fmt.Sprintf("INSERT INTO likesFromUser(uuidUser, uuidLiked, likeOrDislike) VALUES('%s', '%s', '%d')", t.USER.Uuid, uuid, newLike)
				}

				db.Exec(query)

				queryCheckLikes := fmt.Sprintf("SELECT likeOrDislike FROM likesFromUser WHERE uuidLiked = '%s'", uuid)
				row, err := db.Query(queryCheckLikes)
				if err != nil {
					fmt.Println(err)
				} else {
					for row.Next() {
						defer row.Close()
						row.Scan(&likeOrDislike)
						allLikes += likeOrDislike
					}

					queryUpdateMsg := fmt.Sprintf("UPDATE messages SET like = '%d' WHERE uuid = '%s'", allLikes, uuid)
					db.Exec(queryUpdateMsg)
				}
			}
		}
	} else {
		fmt.Println("Vous devez être connecté pour aimer ou ne pas aimer un message.")
	}
}
