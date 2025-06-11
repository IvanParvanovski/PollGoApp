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

var allPolls []models.Poll


func GetAllPolls() []models.Poll {
	// 1) Create a fresh context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
	
	// 2) Point to the "poll" collection 
	collection := globals.MONGO_DB.Collection("polls")
	fmt.Println(collection)
	// collection := globals.MONGO_CLIENT.Collection("testing")

	// 3) Execute an unfiltered find
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("GetAllPolls: find error: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	// 4) Decode all documents into a slice
	var allPolls []models.Poll
	if err := cursor.All(ctx, &allPolls); err != nil {
		log.Printf("GetAllPolls: decode error: %v", err)
		return nil
	}

	return allPolls
}

func findPollIndexById(id primitive.ObjectID) int {
	for i, poll := range allPolls {
		if poll.Id == id {
			return i
		}
	}

	return -1
}

// Poll creation and modification

func AddPoll(poll models.Poll) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	collection := globals.MONGO_DB.Collection("polls")

	// generate a new ObjectID if none provided
	if poll.Id.IsZero() {
		poll.Id = primitive.NewObjectID()
	}

	// check for existing
	count, err := collection.CountDocuments(ctx, bson.M{"_id": poll.Id})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("already exists")
	}

	_, err = collection.InsertOne(ctx, poll)
	
	return err
}

func RemovePollById(id primitive.ObjectID) error {
	pollIndex := findPollIndexById(id)
	
	if pollIndex == -1 {
		return errors.New("poll not found")
	}

	allPolls = append(allPolls[:pollIndex], allPolls[pollIndex + 1:]...)
	return nil
}

func EditPollById(id primitive.ObjectID, title string) (models.Poll, error) {
	index := findPollIndexById(id)

	if index == -1 {
		return models.Poll{}, errors.New("poll not found")
	}

	currentPoll := &allPolls[index]
	currentPoll.Title = title

	return *currentPoll, nil 
}

