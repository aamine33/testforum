package forum

import (
	"database/sql"
	"fmt"
	t "forum/listTopics"
	t3 "forum/structs"
	t2 "forum/users"
	t4 "forum/views"
	"net/http"
	"strings"
)

var MESSAGES t3.Messages

func MessagesPageDisplay(db *sql.DB, r *http.Request) {
	MESSAGES.SessionUser = t2.USER.Username

	var creationDate string
	var owner string
	var report int
	var uuid string
	var likeOrDislike int
	var message string
	var id int
	var like int
	var filter string
	var edited int
	var uuidPath string
	var ascDesc string
	var admin int

	ascDesc = "DESC"
	filter = r.FormValue("filter")
	if filter == "" {
		filter = "creationDate"
	} else if filter == "oldest" {
		filter = "creationDate"
		ascDesc = "ASC"
	}

	topicName := strings.Split(r.URL.Path, "/")
	queryTopicName := fmt.Sprintf("SELECT uuid FROM topics WHERE name = '%s'", topicName[2])
	row, err := db.Query(queryTopicName)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&uuid)
			if err != nil {
				fmt.Println(err)
			}
		}
		row.Close()
		uuidPath = uuid
	}

	queryNew := fmt.Sprintf("SELECT uuid FROM users WHERE username = '%s'", t2.USER.Username)
	row, err = db.Query(queryNew)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&uuid)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	for i := 0; i < len(t.TOPICSANDSESSION.Topics); i++ {
		MESSAGES.CreationDate = t.TOPICSANDSESSION.Topics[i].CreationDate
		MESSAGES.Name = t.TOPICSANDSESSION.Topics[i].Name
		MESSAGES.Owner = t.TOPICSANDSESSION.Topics[i].Owner
		MESSAGES.Likes = t.TOPICSANDSESSION.Topics[i].Likes
		MESSAGES.Id = t.TOPICSANDSESSION.Topics[i].Id
		MESSAGES.UuidPath = t.TOPICSANDSESSION.Topics[i].Uuid
	}

	query := fmt.Sprintf("SELECT id, message, creationDate, owner, report, like, edited, uuid FROM messages WHERE uuidPath = '%s' ORDER BY %s %s ", uuidPath, filter, ascDesc)
	row, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		MESSAGES.Messages = nil
		for row.Next() {
			err = row.Scan(&id, &message, &creationDate, &owner, &report, &like, &edited, &uuid)
			if err != nil {
				fmt.Println(err)
			} else {
				messageIndex := len(MESSAGES.Messages)

				// Ajout du message à la liste des messages
				MESSAGES.Messages = append(MESSAGES.Messages, t3.Message{})
				MESSAGES.Messages[messageIndex].Id = id
				MESSAGES.Messages[messageIndex].Message = message
				MESSAGES.Messages[messageIndex].CreationDate = t4.DisplayTime(creationDate)
				MESSAGES.Messages[messageIndex].Owner = owner
				MESSAGES.Messages[messageIndex].Report = report
				MESSAGES.Messages[messageIndex].Uuid = uuid
				MESSAGES.Messages[messageIndex].Like = like
				MESSAGES.Messages[messageIndex].Edited = edited

				if MESSAGES.SessionUser != "" {
					queryGetIfAdmin := fmt.Sprintf("SELECT admin FROM users WHERE username = '%s'", MESSAGES.SessionUser)
					row, err := db.Query(queryGetIfAdmin)
					if err != nil {
						fmt.Println(err)
					} else {
						for row.Next() {
							err = row.Scan(&admin)
							if err != nil {
								fmt.Println(err)
							}
						}
					}

					if owner == MESSAGES.SessionUser || admin == 1 {
						MESSAGES.Messages[messageIndex].IsOwnerOrAdmin = 1
					}

					checkIfLiked := fmt.Sprintf("SELECT likeOrDislike FROM likesFromUser WHERE uuidUser = '%s' AND uuidLiked = '%s'", uuidUserSession, uuid)
					row, err = db.Query(checkIfLiked)
					if err != nil {
						fmt.Println(err)
					} else {
						for row.Next() {
							err = row.Scan(&likeOrDislike)
							if err != nil {
								fmt.Println(err)
							} else {
								if likeOrDislike == 1 {
									MESSAGES.Messages[messageIndex].IsLiked = 1
								} else if likeOrDislike == -1 {
									MESSAGES.Messages[messageIndex].IsDisliked = 1
								}
							}
						}
					}
				}
			}
		}
	}
}
