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

var _ (PersistenceInterface) = (*PersistenceImplementation)(nil)

type UserID string

type PersistenceInterface interface {
	Initialize(ctx context.Context) error

	Destroy(ctx context.Context) error

	CreateUser(ctx context.Context, user User) error

	FindUserByUsername(ctx context.Context, username string) (User, error)

	DeleteUser(ctx context.Context) error

	UpdateUser(ctx context.Context) error
}

type PersistenceImplementation struct {
	cfg *configuration.PersistenceConfiguration
}

func NewPersistenceImplementation() *PersistenceImplementation {
	return &PersistenceImplementation{
		cfg: &configuration.PersistenceConfiguration{Hostname: "10.0.0.77", Port: 27017, Username: "root", Password: "example"},
	}
}

func NewPersistenceImplementationConfiguration(cfg *configuration.PersistenceConfiguration) *PersistenceImplementation {
	return &PersistenceImplementation{
		cfg: cfg,
	}
}

func (impl *PersistenceImplementation) CreateUser(ctx context.Context, user User) (UserID, error) {
	p, err := persistence.NewPersistence(ctx, impl.cfg, logging.NewLogger())
	if err != nil {
		return "", err
	}
	result, err = p.Create(ctx, databaseName, userCollectionName, user)
	if err != nil {
		return "", err
	}
	return UseriD(result.InsertedID.String()), nil
}

func (impl *PersistenceImplementation) FindUserByUsername(ctx context.Context, username string) (User, error) {
	p, err := persistence.NewPersistence(ctx, impl.cfg, logging.NewLogger())
	if err != nil {
		return "", err
	}
	var user models.UserModel
	p.Find(ctx, databaseName, userCollectionName, getUserByUsernameFilter, user)
	return nil
}

func (impl *PersistenceImplementation) DeleteUser(ctx context.Context) error {
	return nil
}

func (impl *PersistenceImplementation) UpdateUser(ctx context.Context) error {
	return nil
}

// Database Name

var databaseName = "user"

var database = persistence.NewPDatabase(databaseName)

func (impl *PersistenceImplementation) Initialize(ctx context.Context) error {

	p, err := persistence.NewPersistence(ctx, impl.cfg, logging.NewLogger())
	if err != nil {
		return err
	}

	database.Options = options.DatabaseOptions{}

	userCollection.AddIndex(usernameIndexName, usernameIndex)
	userCollection.AddIndex(usernamePasswordIndexName, usernamePasswordIndex)

	userCollection.AddValidator(userValidatorName, userValidator)

	database.AddCollection(ctx, userCollectionName, *userCollection)

	p.AddDatabase(ctx, databaseName, *database)

	err = p.Generate(ctx)
	if err != nil {
		return err
	}

	// userCollection.AddFilter()

	// catalogCollection.AddIndex()
	// catalogCollection.AddValidator()
	// catalogCollection.AddFilter()

	return nil
}

func (impl *PersistenceImplementation) Destroy(ctx context.Context) error {
	p, err := persistence.NewPersistence(ctx, impl.cfg, logging.NewLogger())
	if err != nil {
		return err
	}
	p.Drop(ctx, databaseName)
	return nil
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

var getUserByUsernameFilterName = "GetUserByUsername"

func getUserByUsernameFilter(username string) bson.M {
	return bson.M{
		"username": username,
	}
}
