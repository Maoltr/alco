package http

import (
	"github.com/Maoltr/alco/internal/api/beer/service"
	"github.com/Maoltr/alco/internal/api/beer/transport"
	"github.com/Maoltr/alco/pkg/config"
	"github.com/Maoltr/alco/pkg/logger"
	"github.com/Maoltr/alco/pkg/server"
	"github.com/labstack/echo"
)

func Start(config config.Server, beerSvc service.Service, logger logger.Logger) {
	e := echo.New()

	transport.NewService(e, beerSvc, logger)

	server.Start(e, config)
}
