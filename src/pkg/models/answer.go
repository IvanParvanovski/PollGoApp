package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Answer struct {
	Id          primitive.ObjectID  `bson:"_id" json:"id"`
	Description string	  		    `bson:"description" json:"description"`
}
