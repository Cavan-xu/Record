package common

import "time"

func GetCurMillionSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeMillionSeconds(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
