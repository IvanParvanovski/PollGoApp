package models 

type UserResponse struct {
	Id           uint32
	SubmissionId uint32
	QuestionId   uint32
	AnswerIds    []uint32
	PollId       uint32
}

