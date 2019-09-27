package server

import (
	"github.com/Maoltr/alco/pkg/config"
	"github.com/labstack/echo"
	"time"
)

// Start starts echo server
func Start(e *echo.Echo, cfg config.Server) {
	e.Server.Addr = cfg.Port
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeoutSeconds) * time.Second
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeoutSeconds) * time.Second
	e.Debug = cfg.Debug
	e.Logger.Fatal(e.Start(cfg.Port))
}
