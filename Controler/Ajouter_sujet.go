package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func AddTopic(r *http.Request, database *sql.DB) {
	if r.Method != http.MethodPost {
		return
	}

	category := strings.Split(r.URL.Path, "/")
	category = strings.Split(category[2], "=")

	topicName := r.FormValue("topic_name")
	firstMessage := r.FormValue("firstMessage")

	if len(topicName) < 4 {
		fmt.Println("Le nom du sujet doit contenir au moins 4 caractères.")
		return
	}

	if users.USER.Username == "" {
		fmt.Println("Vous devez être connecté pour créer un sujet.")
		return
	}

	query := `SELECT name FROM topics WHERE name = ?`
	row := database.QueryRow(query, topicName)
	var existingName string
	err := row.Scan(&existingName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Erreur lors de la vérification de l'unicité du nom du sujet :", err)
		return
	}
	if existingName != "" {
		fmt.Println("Le nom du sujet est déjà pris.")
		return
	}

	creationDate := time.Now().String()
	uuid := uuid.New()
	query = `INSERT INTO topics(name, firstMessage, creationDate, owner, likes, nmbPosts, lastPost, category, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = database.Exec(query, topicName, firstMessage, creationDate, users.USER.Username, 0, 0, "0", category[1], uuid)
	if err != nil {
		fmt.Println("Erreur lors de la création du sujet :", err)
		return
	}

	fmt.Println("Nouveau sujet ajouté :", topicName)

	if len(firstMessage) >= 2 {
		err = AddFirstMessageInMessages(firstMessage, creationDate, users.USER.Username, uuid, database)
		if err != nil {
			fmt.Println("Erreur lors de l'ajout du premier message :", err)
			return
		}
	} else {
		fmt.Println("Le premier message doit contenir au moins 2 caractères.")
	}
}
