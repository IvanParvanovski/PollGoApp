package models

import (
	"github.com/google/uuid"
)


type Answer struct {
	Id          uuid.UUID
	Description string
}