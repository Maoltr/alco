package main

import (
	"context"
	"fmt"
	"github.com/Maoltr/alco/domain"
	"github.com/Maoltr/alco/external/mongo"
	"github.com/Maoltr/alco/internal/api/repositories"
	"github.com/Maoltr/alco/pkg/config"
	"github.com/Maoltr/alco/pkg/logger"
	"github.com/satori/go.uuid"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := config.NewConfig("./cmd/api/config.json")
	if err != nil {
		panic(fmt.Sprintf("can not parse config file, message:%s", err.Error()))
	}

	logger := logger.New(config.Logger)

	mongoClient, err := mongo.NewConnectionWithoutChecks(ctx, config.Mongo)
	if err != nil {
		panic(err.Error())
	}
	defer mongoClient.Disconnect(ctx)

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database(config.Mongo.DatabaseName)

	collection := db.Collection(config.Mongo.Collections.Beer)

	beerRepo := repositories.NewBeer(collection, logger)

	beer := domain.Beer{ID: uuid.NewV4().String(), Name: "Test"}
	if err := beerRepo.Create(ctx, beer); err != nil {
		panic(err)
	}

	if err := beerRepo.Delete(ctx, beer.ID); err != nil {
		panic(err)
	}
}
