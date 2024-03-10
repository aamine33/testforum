package forum

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AddFirstMessageInMessages(firstMessage string, creationDate time.Time, owner string, uuidPath uuid.UUID, db *sql.DB) error {
	uuid := uuid.New()
	addInMessages := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, edited, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	query, err := db.Prepare(addInMessages)
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête d'insertion de message: %v", err)
	}
	defer query.Close()

	_, err = query.Exec(firstMessage, creationDate, owner, 0, uuidPath, 0, 0, uuid)
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de la requête d'insertion de message: %v", err)
	}
	return nil
}
