package utils

import (
	"fmt"
	"time"
)

func GetDateNowFormated() string {
	now := time.Now()

	requestDate := fmt.Sprintf(
		"%02d/%02d/%d %02d:%02d:%02d",
		now.Day(),
		now.Month(),
		now.Year(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)

	return requestDate
}
