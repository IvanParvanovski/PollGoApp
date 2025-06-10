package models

import "github.com/google/uuid"

type UserResponse struct {
	Id         uuid.UUID	`json:"id"`
	AnswerIds  []uuid.UUID	`json:"answerIds"`
	PollId     uuid.UUID	`json:"pollId"`
}

