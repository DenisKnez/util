package util

import (
	"github.com/gofrs/uuid"
)

type UUID struct {
	uuid *uuid.UUID
}

func (u *UUID) GetNewUUID() (uuid.UUID, error) {
	if u.uuid != nil {
		return *u.uuid, nil
	}

	return uuid.NewV4()
}
