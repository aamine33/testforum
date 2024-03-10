package forum 

import (
	"fmt"
	"net/http"
)

type Cookies struct {
	UuidUser string
}

var COOKIES Cookies

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    USER.Uuid,
		Path:     "/",
		MaxAge:   99999999999,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	fmt.Println("Cookie set")
}
