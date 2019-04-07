package models

import "time"

func toMS(t time.Time) int64 {
	return t.UTC().UnixNano() / int64(time.Millisecond)
}
