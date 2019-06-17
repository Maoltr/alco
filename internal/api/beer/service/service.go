package service

import (
	"context"
	"github.com/Maoltr/alco/internal/api/repositories"
	"github.com/Maoltr/alco/internal/pkg/structs"
	"github.com/Maoltr/alco/model"
	"github.com/Maoltr/alco/pkg/extendedError"
	"github.com/Maoltr/alco/pkg/logger"
	"net/http"
)

// Service represents interface for beer service
type Service interface {
	Create(ctx context.Context, req CreateBeerRequest) (model.Beer, error)
	Get(ctx context.Context, id string) (model.Beer, error)
	List(ctx context.Context) ([]model.Beer, error)
	Update(ctx context.Context, req UpdateBeerRequest) (model.Beer, error)
	Delete(ctx context.Context, id string) error
}

// NewBeerService returns instance of beer service
func NewBeerService(beerRepository repositories.Beer, logger logger.Logger) Service {
	return &service{beerRepository: beerRepository, logger: logger}
}

type service struct {
	beerRepository repositories.Beer
	logger         logger.Logger
}

// NewBeer creates new beer
func (s *service) Create(ctx context.Context, req CreateBeerRequest) (model.Beer, error) {
	beer, err := req.ConvertToBeer()
	if err != nil {
		return model.Beer{}, extendedError.NewWithStatus(http.StatusBadRequest, err.Error())
	}

	if err := s.beerRepository.Create(ctx, beer); err != nil {
		return model.Beer{}, extendedError.New(err.Error())
	}

	return beer, nil
}

func (s *service) Get(ctx context.Context, id string) (model.Beer, error) {
	beer, err := s.beerRepository.Get(ctx, id)
	if err != nil {
		return model.Beer{}, extendedError.New(err.Error())
	}

	return beer, err
}

func (s *service) List(ctx context.Context) ([]model.Beer, error) {
	beers, err := s.beerRepository.List(ctx)
	if err != nil {
		return nil, extendedError.New(err.Error())
	}

	return beers, nil
}

func (s *service) Update(ctx context.Context, req UpdateBeerRequest) (model.Beer, error) {
	err := req.IsValid()
	if err != nil {
		return model.Beer{}, extendedError.NewWithStatus(http.StatusBadRequest, err.Error())
	}

	beer, err := s.beerRepository.Get(ctx, req.ID)
	if err != nil {
		return model.Beer{}, extendedError.NewWithStatus(http.StatusNotFound, err.Error())
	}

	structs.Merge(&beer, &req)

	err = s.beerRepository.Update(ctx, beer)
	if err != nil {
		return model.Beer{}, extendedError.New(err.Error())
	}

	return beer, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.beerRepository.Delete(ctx, id); err != nil {
		return extendedError.New(err.Error())
	}

	return nil
}
