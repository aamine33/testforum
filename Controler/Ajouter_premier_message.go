package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
	t2 "forum/views"
	"net/http"
)

var TOPICSANDSESSION t.TopicsAndSession

func GetRandomMessages(db *sql.DB, r *http.Request) {
	var name string
	var firstMessage string
	var creationDate string
	var owner string
	var likes int
	// var nmbPosts int
	var category string
	var allTopics []string
	var alreadyUsed bool
	var username string

	TOPICSANDSESSION.Topics = nil
	allTopics = nil

	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		TOPICSANDSESSION.SessionUser = ""
	} else {
		queryGetName := fmt.Sprintf("SELECT username FROM users WHERE uuid = '%s'", cookie.Value)
		row, err := db.Query(queryGetName)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				row.Scan(&username)
			}
			TOPICSANDSESSION.SessionUser = username
		}
	}

	//getting this amount of random topic
	for i := 0; i < 10; i++ {
		query := "SELECT name, firstMessage, creationDate, owner, likes, category FROM topics ORDER BY RANDOM() LIMIT 1"
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				row.Scan(&name, &firstMessage, &creationDate, &owner, &likes, &category)
			}
			alreadyUsed = false
			for i := 0; i < len(allTopics); i++ {
				if allTopics[i] == name {
					alreadyUsed = true
				}
			}
			if !alreadyUsed {
				allTopics = append(allTopics, name)
				topicIndex := len(TOPICSANDSESSION.Topics)
				TOPICSANDSESSION.Topics = append(TOPICSANDSESSION.Topics, t.Topic{})
				TOPICSANDSESSION.Topics[topicIndex].Name = name
				TOPICSANDSESSION.Topics[topicIndex].Likes = likes
				TOPICSANDSESSION.Topics[topicIndex].CreationDate = t2.DisplayTime(creationDate)
				TOPICSANDSESSION.Topics[topicIndex].Owner = owner
				TOPICSANDSESSION.Topics[topicIndex].FirstMessage = firstMessage
				TOPICSANDSESSION.Topics[topicIndex].Category = category
				// TOPICSANDSESSION.Topics[topicIndex].NmbPosts = nmbPosts
			}
		}
	}
}
