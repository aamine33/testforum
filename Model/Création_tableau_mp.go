package forum

import (
	"database/sql"
	"fmt"
)

const (
	colUser1        = "user1"
	colUser2        = "user2"
	colCreationDate = "creationDate"
	colMessage      = "message"
)

func CreateTableMp(db *sql.DB) {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS mp(
		"%s" TEXT,
		"%s" TEXT,
		"%s" TEXT,
		"%s" TEXT);`, colUser1, colUser2, colCreationDate, colMessage)

	query, err := db.Prepare(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = query.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table mp created successfully")
}
