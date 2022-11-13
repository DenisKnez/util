package util

import (
	"time"

	"github.com/gofrs/uuid"
)

type Util struct {
	time       *Time
	uuid       *UUID
	stringUtil *String
}

//New
// now take the current time or nil
// uuid takes a random uuid or nil
func New(now *time.Time, uuid *uuid.UUID) *Util {
	return &Util{
		time: &Time{
			now: now,
		},
		uuid: &UUID{
			uuid: uuid,
		},
		stringUtil: &String{},
	}
}

func (u *Util) String() *String {
	return u.stringUtil
}

func (u *Util) Time() *Time {
	return u.time
}

func (u *Util) UUID() *UUID {
	return u.uuid
}
