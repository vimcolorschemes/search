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
	setAuth(clientOptions)

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

func setAuth(clientOptions *options.ClientOptions) {
	databaseUsername, usernameExists := dotenv.Get("MONGODB_USERNAME")
	databasePassword, passwordExists := dotenv.Get("MONGODB_PASSWORD")
	if usernameExists && databaseUsername != "" && passwordExists && databasePassword != "" {
		credentials := options.Credential{Username: databaseUsername, Password: databasePassword}
		clientOptions.SetAuth(credentials)
	}
}

// StoreSearchIndex stores the payload in the search index collection
func StoreSearchIndex(searchIndex interface{}) bool {
	return false
}

// GetSearchIndex stores the payload in the search index collection
func GetSearchIndex() interface{} {
	var result map[string]interface{}

	err := searchIndexCollection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result["index"]
}
