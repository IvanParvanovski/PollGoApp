package models

type QuestionInterface interface {
	GetQuestion() Question
}

func (m MultipleChoice) GetQuestion() Question {
	return m.Question
}
func (s SingleChoice) GetQuestion() Question {
	return s.Question
}

type Question struct {
	Id              uint32
	Description     string
	Required        bool
	PossibleAnswers []Answer
}

type MultipleChoice struct {
	Question

	MinSelection int
	MaxSelection int
}
type SingleChoice struct {
	Question
}