package database

import (
	"context"
	"log"
	"strings"

	"github.com/vimcolorschemes/search/internal/dotenv"
	"github.com/vimcolorschemes/search/internal/repository"
	"github.com/vimcolorschemes/search/internal/request"
	str "github.com/vimcolorschemes/search/internal/string"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var searchIndexCollection *mongo.Collection

func init() {
	connectionString, exists := dotenv.Get("MONGODB_CONNECTION_STRING")
	if !exists {
		log.Fatal("Database connection string not found in env")
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
func Store(searchIndex []repository.Repository) error {
	deleteResult, err := searchIndexCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error while deleting previous search index")
		return err
	}

	log.Printf("Deleted %d repositories from search index", deleteResult.DeletedCount)

	documents := []interface{}{}

	for _, repository := range searchIndex {
		documents = append(documents, repository)
	}

	insertResult, err := searchIndexCollection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal("Error while inserting new search index")
		return err
	}

	log.Printf("Inserted %d repositories into search index", len(insertResult.InsertedIDs))

	return nil
}

// Search queries the mongo database and returns the result
func Search(parameters request.SearchParameters) ([]repository.Repository, int64, error) {
	filters := bson.D{{Key: "$and", Value: buildSearchQueries(parameters.Query)}}

	options := new(options.FindOptions)
	options.SetSkip(int64((parameters.Page - 1) * parameters.PerPage))
	options.SetLimit(int64(parameters.PerPage))

	cursor, err := searchIndexCollection.Find(ctx, filters, options)
	if err != nil {
		return []repository.Repository{}, -1, err
	}

	defer cursor.Close(ctx)

	var results = []repository.Repository{}
	if err = cursor.All(ctx, &results); err != nil {
		return []repository.Repository{}, -1, err
	}

	total, err := searchIndexCollection.CountDocuments(ctx, filters)
	if err != nil {
		return []repository.Repository{}, -1, err
	}

	return results, total, nil
}

func buildSearchQueries(query string) bson.A {
	queries := bson.A{}

	for _, word := range strings.Split(query, " ") {
		word = str.Normalize(word)
		queries = append(queries,
			bson.D{
				{
					Key: "$or",
					Value: bson.A{
						bson.D{{Key: "name", Value: primitive.Regex{Pattern: word, Options: "i"}}},
						bson.D{{Key: "owner.name", Value: primitive.Regex{Pattern: word, Options: "i"}}},
						bson.D{{Key: "description", Value: primitive.Regex{Pattern: word, Options: "i"}}},
					},
				},
			},
		)
	}

	return queries
}
