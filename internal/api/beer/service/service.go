package service

import (
	"context"
	"github.com/Maoltr/alco/domain"
	"github.com/Maoltr/alco/internal/pkg/structs"
	"github.com/Maoltr/alco/pkg/extendedError"
	"github.com/Maoltr/alco/pkg/logger"
	"net/http"
)

// Service represents interface for beer service
type Service interface {
	Create(ctx context.Context, req CreateBeerRequest) (domain.Beer, error)
	Get(ctx context.Context, id string) (domain.Beer, error)
	List(ctx context.Context) ([]domain.Beer, error)
	Update(ctx context.Context, req UpdateBeerRequest) (domain.Beer, error)
	Delete(ctx context.Context, id string) error
}

// NewBeerService returns instance of beer service
func NewBeerService(Beer domain.BeerRepository, logger logger.Logger) Service {
	return &service{Beer: Beer, logger: logger}
}

type service struct {
	Beer   domain.BeerRepository
	logger logger.Logger
}

// NewBeer creates new beer
func (s *service) Create(ctx context.Context, req CreateBeerRequest) (domain.Beer, error) {
	beer, err := req.ConvertToBeer()
	if err != nil {
		return domain.Beer{}, extendedError.NewWithStatus(http.StatusBadRequest, err.Error())
	}

	if err := s.Beer.Create(ctx, beer); err != nil {
		return domain.Beer{}, extendedError.New(err.Error())
	}

	return beer, nil
}

func (s *service) Get(ctx context.Context, id string) (domain.Beer, error) {
	beer, err := s.Beer.Get(ctx, id)
	if err != nil {
		return domain.Beer{}, extendedError.New(err.Error())
	}

	return beer, err
}

func (s *service) List(ctx context.Context) ([]domain.Beer, error) {
	beers, err := s.Beer.List(ctx)
	if err != nil {
		return nil, extendedError.New(err.Error())
	}

	return beers, nil
}

func (s *service) Update(ctx context.Context, req UpdateBeerRequest) (domain.Beer, error) {
	err := req.IsValid()
	if err != nil {
		return domain.Beer{}, extendedError.NewWithStatus(http.StatusBadRequest, err.Error())
	}

	beer, err := s.Beer.Get(ctx, req.ID)
	if err != nil {
		return domain.Beer{}, extendedError.NewWithStatus(http.StatusNotFound, err.Error())
	}

	structs.Merge(&beer, &req)

	err = s.Beer.Update(ctx, beer)
	if err != nil {
		return domain.Beer{}, extendedError.New(err.Error())
	}

	return beer, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.Beer.Delete(ctx, id); err != nil {
		return extendedError.New(err.Error())
	}

	return nil
}
