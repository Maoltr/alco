package transport

import (
	"context"
	beer "github.com/Maoltr/alco/internal/api/beer/service"
	"github.com/Maoltr/alco/pkg/extendedError"
	"github.com/Maoltr/alco/pkg/logger"
	"github.com/Maoltr/alco/pkg/request"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
)

func NewService(e echo.Echo, beerService beer.Service, logger logger.Logger) {
	beerHTTPSvc := service{svc: beerService, logger: logger}

	e.Group("beers")

	e.POST("", beerHTTPSvc.create)
	e.PATCH("/:id", beerHTTPSvc.update)
	e.GET("", beerHTTPSvc.list)
	e.GET("/:id", beerHTTPSvc.get)
	e.DELETE("/:id", beerHTTPSvc.delete)
}

type service struct {
	svc    beer.Service
	logger logger.Logger
}

func (s *service) create(c echo.Context) error {
	req, err := BeerCreateRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO add users and auth middleware
	serviceReq := req.ConvertToServiceReq(uuid.NewV4().String())

	// TODO add context with timeout
	beer, err := s.svc.Create(context.TODO(), serviceReq)
	if err != nil {
		return c.JSON(err.(extendedError.Error).Status, err.Error())
	}

	return c.JSON(http.StatusOK, beer)
}

func (s *service) update(c echo.Context) error {
	req, err := BeerUpdateRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := request.ID(c)
	if err != nil {
		return err
	}

	serviceReq := req.ConvertToServiceReq(id)

	// TODO add context with timeout
	beer, err := s.svc.Update(context.TODO(), serviceReq)
	if err != nil {
		return c.JSON(err.(extendedError.Error).Status, err.Error())
	}

	return c.JSON(http.StatusOK, beer)
}

func (s *service) get(c echo.Context) error {
	id, err := request.ID(c)
	if err != nil {
		return err
	}

	// TODO add context with timeout
	beer, err := s.svc.Get(context.TODO(), id)
	if err != nil {
		return c.JSON(err.(extendedError.Error).Status, err.Error())
	}

	return c.JSON(http.StatusOK, beer)
}

func (s *service) list(c echo.Context) error {
	// TODO add context with timeout
	beers, err := s.svc.List(context.TODO())
	if err != nil {
		return c.JSON(err.(extendedError.Error).Status, err.Error())
	}

	return c.JSON(http.StatusOK, beers)
}

func (s *service) delete(c echo.Context) error {
	id, err := request.ID(c)
	if err != nil {
		return err
	}

	// TODO add context with timeout
	if err := s.svc.Delete(context.TODO(), id); err != nil {
		return c.JSON(err.(extendedError.Error).Status, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
