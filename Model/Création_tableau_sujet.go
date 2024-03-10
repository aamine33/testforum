package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableTopics(db *sql.DB) error {
	topicTable := `CREATE TABLE IF NOT EXISTS topics (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		firstMessage TEXT,
		creationDate TEXT,
		owner TEXT,
		likes INTEGER,
		nmbPosts INTEGER,
		lastPost TEXT,
		category TEXT,
		uuid TEXT
	);`

	_, err := db.Exec(topicTable)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la table topics : %v", err)
	}

	fmt.Println("Table topics créée avec succès")
	return nil
}
