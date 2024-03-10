package forum

import (
	"strconv"
)

func Atoi(date string) (int, error) {
	marks, err := strconv.Atoi(date)
	if err != nil {
		return 0, err
	}
	return marks, nil
}
