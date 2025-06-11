package models

import ("go.mongodb.org/mongo-driver/bson/primitive")

type UserResponse struct {
	Id        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	AnswerIds []primitive.ObjectID `bson:"answerIds" json:"answerIds"`
	PollId    primitive.ObjectID   `bson:"pollId" json:"pollId"`
}
