package models

import "github.com/google/uuid"

type QuestionType string

const (
	SingleChoiceType   QuestionType = "single"
	MultipleChoiceType QuestionType = "multiple"
)

type Question struct {
	Id              uuid.UUID    `json:"id"`
	Description     string       `json:"description"`
	PossibleAnswers []Answer     `json:"possibleAnswers"`
	QuestionType    QuestionType `json:"questionType"`

	MultipleChoice *MultipleChoice
}

type MultipleChoice struct {
	MinSelection 	int 		`json:"minSelection"`
	MaxSelection 	int 		`json:"maxSelection"`
}
