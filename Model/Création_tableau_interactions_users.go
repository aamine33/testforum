package forum

import (
	"database/sql"
	"fmt"
)

const (
	TextColumn    = "TEXT"
	IntegerColumn = "INTEGER"
)

func CreateTableLikesFromUser(db *sql.DB) error {
	likesFromUserTable := `
	CREATE TABLE IF NOT EXISTS likesFromUser(
		uuidUser TEXT,
		uuidLiked TEXT,
		likeOrDislike INTEGER
	);
	`

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if _, err := tx.Exec(likesFromUserTable); err != nil {
		return fmt.Errorf("failed to create table likesFromUser: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Println("table likesFromUser created successfully")
	return nil
}
