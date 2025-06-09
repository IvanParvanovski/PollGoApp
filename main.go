package main 

type Poll struct {
	Id uint32
	Title string 
	Questions []QuestionInterface
}

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
	Id uint32 
	Description string 
	Required bool
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

type Answer struct {
	Id uint32
	Description string 
}

type Submission struct {
	Id uint32
}

type UserResponse struct {
	Id uint32
	SubmissionId uint32
	QuestionId uint32
	AnswerIds []uint32
	PollId uint32
}

func main() {

}