package main

import (
	"context"
	"fmt"
	"github.com/Maoltr/alco/domain"
	"github.com/Maoltr/alco/external/mongo"
	"github.com/Maoltr/alco/internal/api/beer/service"
	"github.com/Maoltr/alco/internal/api/repositories"
	"github.com/Maoltr/alco/internal/api/transport/http"
	"github.com/Maoltr/alco/pkg/config"
	"github.com/Maoltr/alco/pkg/logger"
	"github.com/satori/go.uuid"
	"os"
	"time"
)

const (
	configEnv = "PATH_TO_CONFIG"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pathToConfigFile := os.Getenv(configEnv)
	if pathToConfigFile == "" {
		panic(fmt.Sprintf("env variable %s can not be empty", configEnv))
	}

	config, err := config.NewConfig(pathToConfigFile)
	if err != nil {
		panic(fmt.Sprintf("can not parse config file, message:%s", err.Error()))
	}

	logger := logger.New(config.Logger)
	mongoDB, err := mongo.NewConnection(ctx, config.Mongo)
	if err != nil {
		panic(err.Error())
	}
	defer mongoDB.Client().Disconnect(ctx)

	// requiredCollections := []string{config.Mongo.Collections.Beer}
	// isCollectionsPresented, err := mongo.IsCollectionsPresented(ctx, requiredCollections, mongoDB)
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// if !isCollectionsPresented {
	// 	panic(fmt.Sprintf("collections: %v, are not presented in databes: %s", requiredCollections, config.Mongo.DatabaseName))
	// }

	beerCollection := mongoDB.Collection(config.Mongo.Collections.Beer)

	beerRepository := repositories.NewBeer(beerCollection, logger)
	beerSvc := service.NewBeerService(beerRepository, logger)

	http.Start(config.Server, beerSvc, logger)
}
