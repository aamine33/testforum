package forum

import (
	"database/sql"
	"log"

	"forum/listtopics"
	"forum/messages"
	"forum/mp"
	"forum/profil"
	"forum/report"
	"forum/users"
)

func CreateDatabaseTables(db *sql.DB) error {
	if err := users.CreateTable(db); err != nil {
		return err
	}

	if err := listtopics.CreateTable(db); err != nil {
		return err
	}

	if err := messages.CreateTable(db); err != nil {
		return err
	}

	if err := profil.CreateTable(db); err != nil {
		return err
	}

	if err := report.CreateTable(db); err != nil {
		return err
	}

	if err := mp.CreateTable(db); err != nil {
		return err
	}

	return nil
}

func InitializeForumDatabase(db *sql.DB) {
	if err := CreateDatabaseTables(db); err != nil {
		log.Fatalf("erreur lors de la création des tables de la base de données : %v", err)
	}
}
