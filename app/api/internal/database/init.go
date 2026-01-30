package database

import (
	"log"

	"github.com/hwangseonu/paperless.dev/internal/common"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func init() {
	config := common.GetConfig()
	uri := config.MongoURI

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	mongoClient, err = mongo.Connect(opts)

	if err != nil {
		log.Fatalln(err)
	}

	mongoDatabase = mongoClient.Database("paperless")
}
