package models

import "github.com/google/uuid"

type UserResponse struct {
	Id         uuid.UUID
	QuestionId uuid.UUID
	AnswerIds  []uuid.UUID
	PollId     uuid.UUID
}

