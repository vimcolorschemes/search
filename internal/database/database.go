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

// Store stores the payload in the search index collection
func Store(searchIndex []interface{}) error {
	deleteResult, err := searchIndexCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Panic("Error while deleting previous search index")
		return err
	}

	log.Printf("Deleted %d repositories from search index", deleteResult.DeletedCount)

	insertResult, err := searchIndexCollection.InsertMany(ctx, searchIndex)
	if err != nil {
		log.Panic("Error while inserting new search index")
		return err
	}

	log.Printf("Inserted %d repositories into search index", len(insertResult.InsertedIDs))

	return nil
}

// Search queries the mongo database and returns the result
func Search(query string, page int, perPage int) ([]interface{}, int) {
	repositories := []interface{}{}
	return repositories, 0
}
