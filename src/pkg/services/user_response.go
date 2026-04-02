package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	globals "mainapp/pkg/global"
	"mainapp/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var votes []models.UserResponse

func GetAllVotes() []models.UserResponse  {
	// 1) Create a ftesh context 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2) Point to the "votes" collection 
	collection := globals.MONGO_DB.Collection("votes")
	fmt.Println(collection)

	// 3) Execute an unfiltered find
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("GetAllVotes: find error: %v", err)
	}
	defer cursor.Close(ctx)
	
	// 4) Decode all documents into a slice
	var votes []models.UserResponse
	if err := cursor.All(ctx, &votes); err != nil {
		log.Printf("GetAllVotes: decode error: %v", err)
		return nil
	}

	return votes
}

func GetPollVotes(pollId primitive.ObjectID) []models.UserResponse {
	var pollVotes []models.UserResponse

	for i, v := range votes {
		if v.PollId == pollId {
			pollVotes = append(pollVotes, votes[i])
		}
	}

	return pollVotes
}

func AddVote(vote models.UserResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := globals.MONGO_DB.Collection("votes")

	if vote.Id.IsZero() {
		vote.Id = primitive.NewObjectID()
	}
	count, err := collection.CountDocuments(ctx, bson.M{"_id": vote.Id})
	
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("aldready exists")
	}

	_, err = collection.InsertOne(ctx, vote)

	return err
}


func RemoveVoteById(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := globals.MONGO_DB.Collection("votes")

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("vote not found")
	}

	return nil
}
