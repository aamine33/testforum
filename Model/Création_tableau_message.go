package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableMessage(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS messages(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			message TEXT,
			creationDate TEXT,
			owner TEXT,
			report INTEGER,
			uuidPath TEXT,
			like INT,
			edited INT,
			uuid TEXT
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Erreur lors de la création de la table Messages:", err)
	} else {
		fmt.Println("Table Messages créée avec succès")
	}
}
