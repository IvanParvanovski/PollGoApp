package models

import (
	"github.com/google/uuid"
)


type Answer struct {
	Id          uuid.UUID `json:"id"`
	Description string	  `json:"description"`
}
