package models

import "github.com/google/uuid"

type Poll struct {
	Id        uuid.UUID
	Title     string
	Question QuestionInterface
}
