package models

import "github.com/google/uuid"

type UserResponse struct {
	Id         uuid.UUID	`json:"id"`
	QuestionId uuid.UUID	`json:"questionId"`
	AnswerIds  []uuid.UUID	`json:"answerIds"`
	PollId     uuid.UUID	`json:"pollId"`
}

