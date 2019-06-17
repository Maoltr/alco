package transport

import (
	beer "github.com/Maoltr/alco/internal/api/beer/service"
	"github.com/labstack/echo"
)

// CreateBeerRequest holds data for creating beer
type CreateBeerRequest struct {
	Name        string `json:"name" validate:"required,min=6,max=100"`
	Description string `json:"description" validate:"required,min=20,max=1000"`
	Strength    uint   `json:"strength" validate:"required,max=100"`
}

func BeerCreateRequest(c echo.Context) (CreateBeerRequest, error) {
	var req CreateBeerRequest

	if err := c.Bind(&req); err != nil {
		return CreateBeerRequest{}, err
	}

	return req, nil
}

// Converts create beer http request to create beer service req
func (c CreateBeerRequest) ConvertToServiceReq(userID string) beer.CreateBeerRequest {
	return beer.CreateBeerRequest{
		Name:        c.Name,
		Description: c.Description,
		Strength:    c.Strength,
		UserID:      userID,
	}
}

// UpdateBeerRequest holds data for updating beer
type UpdateBeerRequest struct {
	Name        string `json:"name" validate:"required,min=6,max=100"`
	Description string `json:"description" validate:"required,min=20,max=1000"`
	Strength    uint   `json:"strength" validate:"required,max=100"`
}

func BeerUpdateRequest(c echo.Context) (UpdateBeerRequest, error) {
	var req UpdateBeerRequest

	if err := c.Bind(&req); err != nil {
		return UpdateBeerRequest{}, err
	}

	return req, nil
}

// Converts update beer http request to update beer service req
func (c UpdateBeerRequest) ConvertToServiceReq(id string) beer.UpdateBeerRequest {
	return beer.UpdateBeerRequest{
		ID:          id,
		Name:        &c.Name,
		Description: &c.Description,
		Strength:    &c.Strength,
	}
}
