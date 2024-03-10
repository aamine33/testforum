package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
	t2 "forum/views"
	"net/http"
	"strings"
)

func DisplayTopic(r *http.Request, db *sql.DB) (t.TopicsAndSession, error) {
	var TOPICSANDSESSION t.TopicsAndSession

	var category, username string
	filter := r.FormValue("filter")
	if filter == "" {
		filter = "creationDate"
	}
	DESCOASC := "DESC"
	if filter == "oldest" {
		DESCOASC = "ASC"
	}

	categoryURL := strings.Split(r.URL.Path, "/")[2]
	urlCat := strings.Split(categoryURL, "=")[1]

	query := fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, category, uuid FROM topics WHERE category = '%s' ORDER BY %s %s", urlCat, filter, DESCOASC)

	if r.FormValue("searchbar") != "" {
		searchName := "%" + r.FormValue("searchbar") + "%"
		query = fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, category, uuid FROM topics WHERE name LIKE '%s' AND category = '%s' ORDER BY %s %s", searchName, urlCat, filter, DESCOASC)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		return TOPICSANDSESSION, fmt.Errorf("erreur lors de la récupération du cookie de session : %v", err)
	}

	row, err := db.Query(query)
	if err != nil {
		return TOPICSANDSESSION, fmt.Errorf("erreur lors de l'exécution de la requête : %v", err)
	}
	defer row.Close()

	for row.Next() {
		var topic t.Topic
		err = row.Scan(&topic.ID, &topic.Name, &topic.FirstMessage, &topic.CreationDate, &topic.Owner, &topic.Likes, &topic.NmbPosts, &category, &topic.UUID)
		if err != nil {
			return TOPICSANDSESSION, fmt.Errorf("erreur lors du scan des résultats de la requête : %v", err)
		}

		topic.CreationDate = t2.DisplayTime(topic.CreationDate)

		if cookie.Value != "" {
			var admin int
			adminQuery := fmt.Sprintf("SELECT admin FROM users WHERE uuid = '%s'", cookie.Value)
			err = db.QueryRow(adminQuery).Scan(&admin)
			if err != nil {
				return TOPICSANDSESSION, fmt.Errorf("erreur lors de la récupération des informations de l'utilisateur : %v", err)
			}

			if topic.Owner == username || admin == 1 {
				topic.IsOwnerOrAdmin = 1
			}

			likeQuery := fmt.Sprintf("SELECT likeOrDislike FROM likesFromUser WHERE uuidUser = '%s' AND uuidLiked = '%s'", cookie.Value, topic.UUID)
			var likeOrDislike int
			err = db.QueryRow(likeQuery).Scan(&likeOrDislike)
			if err == nil {
				if likeOrDislike == 1 {
					topic.IsLiked = 1
				} else if likeOrDislike == -1 {
					topic.IsDisliked = 1
				}
			}
		}

		TOPICSANDSESSION.Topics = append(TOPICSANDSESSION.Topics, topic)
	}

	return TOPICSANDSESSION, nil
}
