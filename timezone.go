package timezone

import (
	"time"
)

func GetIndonesiaTime() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
}

func GetIndonesiaTimeString() string {
	return time.Now().Format("02 January 2006 15:04:05")
}
