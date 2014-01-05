package app

import (
	"time"
)

func dateformat(unixnano int64) string {
	return time.Unix(0, unixnano).Format("2006-01-02 15:04:05")
}
