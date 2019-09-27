package domain

import "context"

type BeerRepository interface {
	Create(ctx context.Context, beer Beer) error
	Get(ctx context.Context, id string) (Beer, error)
	List(ctx context.Context) ([]Beer, error)
	Update(ctx context.Context, beer Beer) error
	Delete(ctx context.Context, id string) error
}

type Beer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Strength    uint   `json:"strength"`
	AddedBy     string `json:"added_by"`
}
