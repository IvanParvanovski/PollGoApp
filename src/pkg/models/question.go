package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionType string

const (
	SingleChoiceType   QuestionType = "single"
	MultipleChoiceType QuestionType = "multiple"
)

type Question struct {
	Id              primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	Description     string       `bson:"description" json:"description"`
	PossibleAnswers []Answer     `bson:"possibleAnswers" json:"possibleAnswers"`
	QuestionType    QuestionType `bson:"type" json:"type"`

	MultipleChoice *MultipleChoice `bson:"multipleChoice,omitempty" json:"multipleChoice,omitempty"`
}

type MultipleChoice struct {
	MinSelection 	int 		`bson:"minSelection" json:"minSelection"`
	MaxSelection 	int 		`bson:"maxSelection" json:"maxSelection"`
}
