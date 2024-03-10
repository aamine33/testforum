package forum

import (
	"database/sql"
	"log"

)

func AddUsers(db *sql.DB, username string, password string, email string, creationDate string, birthDate string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Erreur en démarrant la transaction :", err)
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", username, email).Scan(&count)
	if err != nil {
		log.Fatal("Erreur en vérifiant l'existence de l'utilisateur ou de l'email :", err)
	}

	if count > 0 {
		if count == 2 {
			log.Println("L'utilisateur et l'email sont déjà pris")
		} else if username == username {
			log.Println("Le nom d'utilisateur est déjà pris")
		} else {
			log.Println("L'email est déjà pris")
		}
		return
	}

	userUUID := userUUID.New()

	_, err = tx.Exec("INSERT INTO users(username, password, email, creationDate, birthDate, admin, reports, uuid, ban) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		username, password, email, creationDate, birthDate, 0, 0, userUUID, 0)
	if err != nil {
		log.Fatal("Erreur en insérant le nouvel utilisateur :", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Erreur en commettant la transaction :", err)
	}

	log.Println("Ajout du nouvel utilisateur :", username)
}
