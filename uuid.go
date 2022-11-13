package util

import (
	"github.com/gofrs/uuid"
)

type UUID struct {
	uuid *uuid.UUID
}

func (u *UUID) GetNewUUID() uuid.UUID {
	if u.uuid != nil {
		return *u.uuid
	}

	return *u.uuid
}
