package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableReports(db *sql.DB) {
	reportsTableQuery := `
	CREATE TABLE IF NOT EXISTS reports (
		uuidUser TEXT,
		uuidReported TEXT
	);`

	stmt, err := db.Prepare(reportsTableQuery)
	if err != nil {
		fmt.Println("Error preparing SQL statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error creating reports table:", err)
		return
	}

	fmt.Println("Table 'reports' created successfully")
}
