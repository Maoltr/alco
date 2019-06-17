package beer

import (
	"context"
	"github.com/Maoltr/alco/external/logger"
	"github.com/Maoltr/alco/internal/api/repositories"
	"github.com/Maoltr/alco/model"
)

// Service represents interface for beer service
type Service interface {
	Create(ctx context.Context, req CreateBeerRequest) (model.Beer, error)
	Get(ctx context.Context, id string) (model.Beer, error)
	List(ctx context.Context) ([]model.Beer, error)
	Update(ctx context.Context, req UpdateBeerRequest) (model.Beer, error)
	Delete(ctx context.Context, id string) error
}

// NewService returns instance of beer service
func NewService(beerRepository repositories.Beer, logger logger.Logger) Service {
	return &service{beerRepository: beerRepository, logger: logger}
}

type service struct {
	beerRepository repositories.Beer
	logger         logger.Logger
}

// NewBeer creates new beer
func (s *service) Create(ctx context.Context, req CreateBeerRequest) (model.Beer, error) {

}

func (s *service) Get(ctx context.Context, id string) (model.Beer, error) {

}

func (s *service) List(ctx context.Context) ([]model.Beer, error) {

}

func (s *service) Update(ctx context.Context, req UpdateBeerRequest) (model.Beer, error) {

}

func (s *service) Delete(ctx context.Context, id string) error {

}

