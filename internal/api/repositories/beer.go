package repositories

import (
	"context"
	"github.com/Maoltr/alco/domain"
	"github.com/Maoltr/alco/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBeer(collection *mongo.Collection, logger logger.Logger) domain.BeerRepository {
	return &beer{collection: collection, logger: logger}
}

type beer struct {
	collection *mongo.Collection
	logger     logger.Logger
}

// Create creates new beer in collection
func (b *beer) Create(ctx context.Context, beer domain.Beer) error {
	result, err := b.collection.InsertOne(ctx, &beer)
	if err != nil {
		b.logger.Errorf("Insert beer error, name: %s, message: %s", beer.Name, err.Error())
		return err
	}

	b.logger.Infof("Inserted beer, id: %s", result.InsertedID)
	return nil
}

// Get returns beer by id
func (b *beer) Get(ctx context.Context, id string) (domain.Beer, error) {
	var result domain.Beer
	filter := bson.D{{Key: "id", Value: id}}

	err := b.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		b.logger.Errorf("Find beer error, id: %s, message: %s", id, err.Error())
		return result, err
	}

	return result, nil
}

// List returns all beers from collection
func (b *beer) List(ctx context.Context) ([]domain.Beer, error) {
	var result []domain.Beer

	cursor, err := b.collection.Find(ctx, bson.D{})
	if err != nil {
		b.logger.Errorf("Get list of beers error, message: %s", err.Error())
		return result, err
	}

	for cursor.Next(ctx) {
		var beer domain.Beer
		if err := cursor.Decode(&beer); err != nil {
			b.logger.Errorf("Decode beer error, message: %s", err.Error())
			return result, err
		}

		result = append(result, beer)
	}

	return result, nil
}

// Update updates beer
func (b *beer) Update(ctx context.Context, beer domain.Beer) error {
	filter := bson.D{{Key: "id", Value: beer.ID}}

	result, err := b.collection.UpdateOne(ctx, filter, &beer)
	if err != nil {
		b.logger.Errorf("Update beer error, name: %s, message: %s", beer.Name, err.Error())
		return err
	}

	b.logger.Infof("Updated beer count", result.ModifiedCount)
	return nil
}

func (b *beer) Delete(ctx context.Context, id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	result, err := b.collection.DeleteOne(ctx, filter)
	if err != nil {
		b.logger.Errorf("Delete beer error, id: %s, message: %s", id, err.Error())
		return err
	}

	b.logger.Infof("Deleted beer count: %v", result.DeletedCount)
	return nil
}
