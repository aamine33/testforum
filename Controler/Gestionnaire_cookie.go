package forum

import (
	"fmt"
	"net/http"
)

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		resetUserData()
	} else {
		if cookie.Value != "" {
			COOKIES.UuidUser = cookie.Value
			LoginWithCookie(cookie.Value)
		} else {
			resetUserData()
		}
	}
}

func resetUserData() {
	USER.Username = ""
	USER.Uuid = ""
	USER.Email = ""
	USER.CreationDate = ""
	USER.Admin = 0
	USER.BirthDate = ""
}
