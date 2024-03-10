package forum

import (
	"fmt"
	"net/http"
)

func CheckCookie(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Erreur lors de la récupération du cookie :", err)
		return false
	}

	COOKIES.UuidUser = cookie.Value
	return true
}
