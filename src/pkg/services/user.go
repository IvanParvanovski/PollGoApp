package services

import (
	"context"
	"errors"
	globals "mainapp/pkg/global"
	"mainapp/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(username, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := globals.MONGO_DB.Collection("users")

	filter := bson.M{"username": username}
	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("username already taken")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Id: primitive.NewObjectID(),
		Username: username, 
		Password: string(hashedPwd),
	}

	_, err = collection.InsertOne(ctx, user)
	return &user, err
}

func FetchUser(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	collection := globals.MONGO_DB.Collection("users")

	
	filter := bson.M{"username": username}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func LogoutUser() error {
	return nil
}

