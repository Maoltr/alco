package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Maoltr/alco/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewConnectionWithoutChecks(ctx context.Context, config config.Mongo) (*mongo.Client, error) {
	// credentials := options.Credential{
	// 	AuthMechanism:           config.Credentials.AuthMechanism,
	// 	AuthMechanismProperties: config.Credentials.AuthMechanismProperties,
	// 	AuthSource:              config.Credentials.AuthSource,
	// 	Username:                config.Credentials.Username,
	// 	Password:                config.Credentials.Password,
	// 	PasswordSet:             config.Credentials.PasswordSet,
	// }

	connectionTimeout := time.Second * time.Duration(config.ConnectionTimeoutInSeconds)
	connectionOpts := options.ClientOptions{
		AppName: &config.AppName,
		//Auth:           &credentials,
		ConnectTimeout: &connectionTimeout,
		Hosts:          config.Hosts,
		MaxPoolSize:    &config.MaxPoolSize,
	}

	client, err := mongo.NewClient(&connectionOpts)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not create mongo db client, config: %v, reason: %s", config, err.Error()))
	}

	if err := client.Connect(ctx); err != nil {
		return nil, errors.New(fmt.Sprintf("can not connect to mongo db, hosts: %s, reason: %s", config.Hosts, err.Error()))
	}

	return client, nil
}

// NewConnection creates connection to mongo database with given options
func NewConnection(ctx context.Context, config config.Mongo) (*mongo.Database, error) {
	conn, err := NewConnectionWithoutChecks(ctx, config)
	if err != nil {
		return nil, err
	}

	databases, err := conn.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not get list of databases, reason: %s", err.Error()))
	}

	var isDBPresented bool
	for _, database := range databases {
		if database == config.DatabaseName {
			isDBPresented = true
			break
		}
	}

	if !isDBPresented {
		return nil, errors.New(fmt.Sprintf("database isn't presented in cluster, db name: %s, presented: %v", config.DatabaseName, databases))
	}

	dbConn := conn.Database(config.DatabaseName)

	return dbConn, nil
}

// IsCollectionsPresented checks is collections with given name presented in mongo database
func IsCollectionsPresented(ctx context.Context, collections []string, db *mongo.Database) (bool, error) {
	cursor, err := db.ListCollections(ctx, bson.D{})
	if err != nil {
		return false, errors.New(fmt.Sprintf("can not get cursor for list of collections, reason: %s", err.Error()))
	}

	requiredCollections := make(map[string]bool)
	for _, collection := range collections {
		requiredCollections[collection] = false
	}

	for cursor.Next(ctx) {
		type collection struct {
			Name string
		}
		coll := collection{}
		if err := cursor.Decode(&coll); err != nil {
			return false, errors.New(fmt.Sprintf("can not decode collection, reason: %s", err.Error()))
		}

		_, ok := requiredCollections[coll.Name]
		if ok {
			requiredCollections[coll.Name] = true
		}
	}

	var notPresentedCollections []string
	for collection, isPresented := range requiredCollections {
		if !isPresented {
			notPresentedCollections = append(notPresentedCollections, collection)
		}
	}

	if len(notPresentedCollections) > 0 {
		return false, errors.New(fmt.Sprintf("given collections isn't presented in databse, names: %v", notPresentedCollections))
	}

	return true, nil
}
