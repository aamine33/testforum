package forum

import (
	"database/sql"
	"fmt"
	"forum/mp"
	"forum/profil"
	"forum/report"
	"html/template"
	"net/http"
	"strings"
)

func Handler_publicProfil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/publicProfil.html"))
	tmpl404 := template.Must(template.ParseFiles("../static/html/404.html"))

	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer databaseForum.Close()

	if r.FormValue("ban") != "" {
		report.Ban(r, databaseForum)
	} else if r.FormValue("report") != "" {
		report.ReportUser(r, databaseForum)
	}

	if r.FormValue("mpMessage") != "" {
		mp.AddMp(r, databaseForum)
	}

	var exists bool
	urlName := strings.Split(r.URL.Path, "/")
	if len(urlName) >= 3 {
		newUrlName := strings.TrimSpace(urlName[2])
		query := "SELECT username FROM users WHERE username = ?"
		row, err := databaseForum.Query(query, newUrlName)
		if err != nil {
			fmt.Println(err)
		} else {
			defer row.Close()
			for row.Next() {
				var name string
				err = row.Scan(&name)
				if err != nil {
					fmt.Println(err)
				} else {
					if name == newUrlName {
						exists = true
					}
				}
			}
		}
	}

	if exists {
		profil.PublicProfil(r, databaseForum)
		tmpl.Execute(w, profil.PUBLICUSER)
	} else {
		tmpl404.Execute(w, nil)
	}
}
