package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
)

func EditTopic(r *http.Request, db *sql.DB) error {
	if r.FormValue("edit") != "" {
		newName := r.FormValue("newName")
		if len(newName) < 2 {
			return fmt.Errorf("le nouveau nom est trop court")
		}

		uuid := strings.Split(r.URL.Path, "/")[2]
		query := fmt.Sprintf("SELECT owner FROM topics WHERE uuid = '%s'", uuid)
		var owner string
		err := db.QueryRow(query).Scan(&owner)
		if err != nil {
			return fmt.Errorf("erreur lors de la récupération du propriétaire du sujet : %v", err)
		}

		if t.USER.Username == owner {
			updateQuery := fmt.Sprintf("UPDATE topics SET name = '%s' WHERE uuid = '%s'", newName, uuid)
			_, err := db.Exec(updateQuery)
			if err != nil {
				return fmt.Errorf("erreur lors de la mise à jour du sujet : %v", err)
			}
			return nil
		}

		return fmt.Errorf("vous devez être le propriétaire du sujet pour le modifier")
	}

	return nil
}
