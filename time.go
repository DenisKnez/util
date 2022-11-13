package util

import "time"

type Time struct {
	now *time.Time
}

func (u *Time) GetCurrentTime() time.Time {
	if u.now != nil {
		return *u.now
	}

	return time.Now()
}
