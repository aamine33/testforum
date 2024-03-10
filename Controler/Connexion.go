package forum

import (
	"database/sql"
	"fmt"
	t4 "forum/listTopics"
	t3 "forum/messages"
	t2 "forum/users"
	t "forum/views"
	"net/http"
)

// Login gère le processus de connexion des utilisateurs
func Login(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	// Informations de connexion
	var email string
	var username string
	var password string
	var birthDate string
	var uuid string
	var creationDate string
	var admin int
	var ban int

	if r.Method == "POST" {
		if t2.USER.Username != "" {
			fmt.Println("Vous êtes déjà connecté")
		} else {
			// Le nom d'utilisateur fonctionnera également avec l'e-mail
			usernameInput := r.FormValue("input_loginusername")
			passwordInput := t.Hash(r.FormValue("input_loginpassword"))

			query := "SELECT username, password, email, creationDate, birthDate, admin, uuid, ban FROM users WHERE username = ? OR email = ?"
			row := db.QueryRow(query, usernameInput, usernameInput)

			err := row.Scan(&username, &password, &email, &creationDate, &birthDate, &admin, &uuid, &ban)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("Aucun utilisateur trouvé avec ces identifiants")
					t4.TOPICSANDSESSION.Error = "Informations d'identification invalides"
					t3.MESSAGES.Error = "Informations d'identification invalides"
				} else {
					fmt.Println("Erreur lors de la recherche de l'utilisateur:", err)
				}
				return
			}

			if ban == 1 {
				fmt.Println("Vous avez été banni, veuillez contacter un administrateur pour lever le bannissement")
				return
			}

			if passwordInput == password {
				fmt.Println("Connexion réussie!")
				// Définir les informations de l'utilisateur connecté
				t2.USER.Username = username
				t2.USER.BirthDate = birthDate
				t2.USER.CreationDate = creationDate
				t2.USER.Email = email
				t2.USER.Uuid = uuid
				t2.USER.Admin = admin

				// Définir le cookie de session pour l'utilisateur connecté
				t2.SetCookieHandler(w, r)
			} else {
				fmt.Println("Mot de passe incorrect")
				t4.TOPICSANDSESSION.Error = "Informations d'identification invalides"
				t3.MESSAGES.Error = "Informations d'identification invalides"
			}
		}
	}
}
