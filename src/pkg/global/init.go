package globals

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	cancelFunc context.CancelFunc
)

func Init(uri string) error {
	MONGO_CONTEXT, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)

	var err error
	MONGO_CLIENT, err = mongo.Connect(
		options.Client().
			ApplyURI(uri).
			SetAuth(options.Credential{
				Username:   "admin",
				Password:   "admin",
				AuthSource: "admin"}))
		
	
	MONGO_DB = MONGO_CLIENT.Database(MONGO_DB_NAME)

	if err != nil {
		cancelFunc()
		return err
	}

	if err = MONGO_CLIENT.Ping(MONGO_CONTEXT, readpref.Primary()); err != nil {
		cancelFunc()
		return err
	}
	fmt.Println("✅ Successfully connected to MongoDB!")

	return nil
}

func DeInit() {

	if err := MONGO_CLIENT.Disconnect(MONGO_CONTEXT); err != nil {
		log.Panicf("failed to disconnect mongo client: %v", err)
	}
}
