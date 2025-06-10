package models

import (
	"github.com/google/uuid"
)

type Poll struct {
	Id       uuid.UUID         `json:"id"`
	Title    string            `json:"title"`
	Question Question 		   `json:"question"`
}

type PollUpdate struct {
	Title string			   `json:"title"`
}
