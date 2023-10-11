package databases

import (
	"log"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	UserCollection *mongo.Collection
)

func Connect()  {
	
}