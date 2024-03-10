package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type PublicUser struct {
	Username      string
	Email         string
	CreationDate  string
	BirthDate     string
	Uuid          string
	Admin         string
	Reports       int
	Ban           int
	TopicsCreated []string
	MessagesSend  []MessageSend
}

type MessageSend struct {
	MessageSendByUser string
	TopicSentInName   string
}

func PublicProfil(r *http.Request, db *sql.DB) {
	namePublic := strings.Split(r.URL.Path, "/")
	username := namePublic[2]

	var publicUser PublicUser

	query := "SELECT username, creationDate, admin, birthDate, reports, ban FROM users WHERE username = ?"
	row := db.QueryRow(query, username)
	err := row.Scan(&publicUser.Username, &publicUser.CreationDate, &publicUser.Admin, &publicUser.BirthDate, &publicUser.Reports, &publicUser.Ban)
	if err != nil {
		fmt.Println(err)
		return
	}

	query = "SELECT message, uuidPath FROM messages WHERE owner = ?"
	rows, err := db.Query(query, publicUser.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var message, uuidPath, topicName string
		err := rows.Scan(&message, &uuidPath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		publicUser.MessagesSend = append(publicUser.MessagesSend, MessageSend{MessageSendByUser: message})

		query := "SELECT name FROM topics WHERE uuid = ?"
		row := db.QueryRow(query, uuidPath)
		err = row.Scan(&topicName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		publicUser.MessagesSend[len(publicUser.MessagesSend)-1].TopicSentInName = topicName
	}

	query = "SELECT name FROM topics WHERE owner = ?"
	rows, err = db.Query(query, publicUser.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var topicName string
		err := rows.Scan(&topicName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		publicUser.TopicsCreated = append(publicUser.TopicsCreated, topicName)
	}

	PUBLICUSER = publicUser
}
