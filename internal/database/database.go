package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/vimcolorschemes/search/internal/dotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var searchIndexCollection *mongo.Collection

func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		// Running in test mode
		return
	}

	connectionString, exists := dotenv.Get("MONGODB_CONNECTION_STRING")
	if !exists {
		log.Panic("Database connection string not found in env")
	}

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	database := client.Database("vimcolorschemes")
	searchIndexCollection = database.Collection("search")
}

// StoreSearchIndex stores the payload in the search index collection
func StoreSearchIndex(searchIndex interface{}) error {
	entry := make(map[string]interface{})
	entry["index"] = searchIndex

	result := searchIndexCollection.FindOneAndReplace(ctx, bson.M{}, entry)
	return result.Err()
}

// GetSearchIndex returns the whole search index
func GetSearchIndex() interface{} {
	var result map[string]interface{}

	err := searchIndexCollection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result["index"]
}
