package forum

import (
	"net/http"
)

func LogOutCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		MaxAge:   -1,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
}
