package globals

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const MONGO_DB_NAME = "pollapp"

var MONGO_DB *mongo.Database
var MONGO_CLIENT *mongo.Client
var MONGO_CONTEXT context.Context
