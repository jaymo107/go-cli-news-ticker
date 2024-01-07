package helpers

import (
	"fmt"
	"time"
)

func GetTimeAgo(unixTime int64) string {
	timestamp := time.Unix(unixTime, 0)
	duration := time.Since(timestamp)

	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24

	var result string

	if days > 0 {
		result = fmt.Sprintf("%d day", days)
		if days > 1 {
			result += "s"
		}

		return result + " ago"
	}

	if hours > 0 {
		if days > 0 {
			result += " and "
		}
		result += fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			result += "s"
		}
	}

	if result == "" {
		return "just now"
	}

	return result + " ago"
}
