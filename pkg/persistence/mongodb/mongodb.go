package mongodb

import (
	"context"

	"github.com/smugalamunga/playground/pkg/configuration"
	"github.com/smugalamunga/playground/pkg/logging"
	"github.com/smugalamunga/playground/pkg/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database Name

var databaseName = "User"

var database = persistence.NewPDatabase(databaseName)

func init() {
	ctx := context.Background()
	p, err := persistence.NewPersistence(ctx, &configuration.PersistenceConfiguration{Hostname: "10.0.0.77", Port: 27017, Username: "root", Password: "example"}, logging.NewLogger())
	if err != nil {
		panic(err)
	}

	database.Options = options.DatabaseOptions{}

	userCollection.AddIndex(usernameIndexName, usernameIndex)
	userCollection.AddIndex(usernamePasswordIndexName, usernamePasswordIndex)

	userCollection.AddValidator(userValidatorName, userValidator)

	database.AddCollection(ctx, userCollectionName, *userCollection)

	p.AddDatabase(ctx, databaseName, *database)

	err = p.Generate(ctx)
	if err != nil {
		panic(err)
	}

	// userCollection.AddFilter()

	// catalogCollection.AddIndex()
	// catalogCollection.AddValidator()
	// catalogCollection.AddFilter()
}

// Collection Names

var userCollectionName = "user"

var userCollection = persistence.NewPCollection(userCollectionName)

var catalogCollectionName = "catalog"

var catalogCollection = persistence.NewPCollection(catalogCollectionName)

// Indexes
var usernameIndexName = "UsernameIdx"
var usernameIndexModel = mongo.IndexModel{
	Keys: bson.M{
		"username": 1,
		// "password": 1,
	},
	Options: options.Index().SetUnique(true).SetName(usernameIndexName),
}
var usernameIndex = persistence.NewIndex(usernameIndexName, usernameIndexModel, options.CreateIndexesOptions{})

var usernamePasswordIndexName = "UsernamePasswordIdx"
var usernamePasswordIndexModel = mongo.IndexModel{
	Keys: bson.D{
		{"username", 1},
		{"password", 1},
	},
	//  error creating index UsernamePasswordIdx, multi-key map passed in for ordered parameter keys
	//Keys: bson.M{
	//	"username": 1,
	//	"password": 1,
	//},
	Options: options.Index().SetUnique(true).SetName(usernamePasswordIndexName),
}
var usernamePasswordIndex = persistence.NewIndex(usernamePasswordIndexName, usernamePasswordIndexModel, options.CreateIndexesOptions{})

// Validators
var userValidatorName = "user"

var userSchema = bson.M{
	"$jsonSchema": bson.M{
		"bsonType": "object",
		"required": []string{"username", "password"},
		"properties": bson.M{
			"username": bson.M{
				"bsonType":    "string",
				"description": "Username for the User",
			},
			"password": bson.M{
				"bsonType":    "string",
				"description": "Password for the User",
			},
		},
	},
}
var userValidator = persistence.NewValidator(userValidatorName, userSchema)

// Filters
