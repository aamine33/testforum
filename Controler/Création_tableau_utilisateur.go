package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableUsers(db *sql.DB) {
	usersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT,
		email TEXT,
		creationDate TEXT,
		birthDate TEXT,
		admin INTEGER,
		reports INTEGER,
		uuid TEXT,
		ban INTEGER
	);`

	_, err := db.Exec(usersTableQuery)
	if err != nil {
		fmt.Println("Erreur lors de la création de la table Users :", err)
		return
	}

	fmt.Println("Table Users créée avec succès")
}
