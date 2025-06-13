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

func GetUserVotes(userId primitive.ObjectID) ([]models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := globals.MONGO_DB.Collection("votes")

	cursor, err := collection.Find(ctx, bson.M{
		"userId": userId,
	})

	if err != nil {
		return []models.UserResponse{}, err
	}

	var votes []models.UserResponse
	if err := cursor.All(ctx, &votes); err != nil {
		return []models.UserResponse{}, err
	}
	
	return votes, nil
}

func contains(id primitive.ObjectID, possibleAnswers []models.Answer) bool {
	for _, a := range possibleAnswers {
		if a.Id == id {
			return true
		}
	}

	return false
}

func AddVote(vote models.UserResponse) (models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if the poll of the vote exists
	pollCollection := globals.MONGO_DB.Collection("polls")
	count, err := pollCollection.CountDocuments(ctx, bson.M{"_id": vote.PollId})

	if err != nil {
		return models.UserResponse{}, err
	}
	if count == 0 {
		return models.UserResponse{}, errors.New("poll with the given id does not exist")
	}

	// Get the poll 
	// check if the answers are valid
	var poll models.Poll
	err = pollCollection.FindOne(ctx, bson.M{"_id": vote.PollId}).Decode(&poll)
	if err != nil {
		return  models.UserResponse{}, err
	}

	for _, requestAnswerId := range vote.AnswerIds {
		if !contains(requestAnswerId, poll.Question.PossibleAnswers) {
			return models.UserResponse{}, errors.New("the given answer id does not exist in the poll")
		}
	} 

	collection := globals.MONGO_DB.Collection("votes")

	// check if the user has already submited a vote to that poll
	submitedVotesCount, err := collection.CountDocuments(ctx, bson.M{"pollId": vote.PollId, "userId": vote.UserId})

	if err != nil {
		return models.UserResponse{}, err
	}
	if submitedVotesCount > 0 {
		return models.UserResponse{}, errors.New("you have already submitted a vote to that poll")
	}

	if vote.Id.IsZero() {
		vote.Id = primitive.NewObjectID()
	}
	count, err = collection.CountDocuments(ctx, bson.M{"_id": vote.Id})
	
	if err != nil {
		return  models.UserResponse{}, err
	}
	if count > 0 {
		return models.UserResponse{}, errors.New("you have voted for this poll already")
	}

	_, err = collection.InsertOne(ctx, vote)

	return vote, err
}


func RemoveVoteById(id, userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() 

	collection := globals.MONGO_DB.Collection("votes")
	
	count, err := collection.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("such a vote does not exist")
	}

	var vote models.UserResponse
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&vote)
	
	if err != nil {
		return err
	}
	if vote.UserId != userId {
		return errors.New("the authenticated user is not an owner of this object")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("vote not found")
	}

	return nil
}
