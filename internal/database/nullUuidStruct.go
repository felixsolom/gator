package database

import "github.com/google/uuid"

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}
