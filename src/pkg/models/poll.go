package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Poll struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId 	 primitive.ObjectID `bson:"userId" json:"userId"`
	Title    string    			`bson:"title" json:"title"`
	Question Question  			`bson:"question" json:"question"`
}

type PollUpdate struct {
	Title string 				`bson:"title,omitempty" json:"title"`
}
