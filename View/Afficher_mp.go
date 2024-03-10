package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	t2 "forum/views"
	"net/http"
	"strings"
)

type MpAndToWho struct {
	ToWho string
	Mps   []Mp
}

type Mp struct {
	PrivateMessage string
	User1          string
	CreationDate   string
}

func DisplayMp(r *http.Request, db *sql.DB) {
	user2 := strings.Split(r.URL.Path, "/")
	toWho := user2[2]

	query := `SELECT user1, message, creationDate FROM mp WHERE (user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?) ORDER BY creationDate ASC`
	rows, err := db.Query(query, t.USER.Username, toWho, toWho, t.USER.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	MPSANDTOWHO := MpAndToWho{
		ToWho: toWho,
		Mps:   make([]Mp, 0),
	}

	for rows.Next() {
		var mp Mp
		err := rows.Scan(&mp.User1, &mp.PrivateMessage, &mp.CreationDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mp.CreationDate = t2.DisplayTime(mp.CreationDate)
		MPSANDTOWHO.Mps = append(MPSANDTOWHO.Mps, mp)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	MPSANDTOWHO = MPSANDTOWHO
}
