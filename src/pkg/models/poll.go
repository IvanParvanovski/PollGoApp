package models

type Poll struct {
	Id        uint32
	Title     string
	Questions []QuestionInterface
}
