package forum

import (
	"fmt"
	"time"
)

func DisplayTime(creationDate string) string {
	if creationDate == "" {
		return ""
	}

	creationTime, err := time.Parse(time.RFC3339, creationDate)
	if err != nil {
		fmt.Println("Error parsing creation date:", err)
		return ""
	}

	now := time.Now()
	diff := now.Sub(creationTime)

	days := int(diff.Hours() / 24)
	months := int(diff.Hours() / 24 / 30)
	years := int(diff.Hours() / 24 / 365)

	if years > 0 {
		if years == 1 {
			return fmt.Sprintf("%d year ago", years)
		}
		return fmt.Sprintf("%d years ago", years)
	} else if months > 0 {
		if months == 1 {
			return fmt.Sprintf("%d month ago", months)
		}
		return fmt.Sprintf("%d months ago", months)
	} else if days > 0 {
		if days == 1 {
			return fmt.Sprintf("%d day ago", days)
		}
		return fmt.Sprintf("%d days ago", days)
	}

	return "Today"
}
