package request

import (
	"github.com/labstack/echo"
	"net/http"
)

const IDLen = 20

var (
	// ErrBadRequest returns status 400
	ErrBadRequest = echo.NewHTTPError(http.StatusBadRequest)
)

// ID returns id url parameter.
// In case of len != IDLen, request will be aborted with StatusBadRequest.
func ID(c echo.Context) (string, error) {
	id := c.Param("id")
	if len(id) != IDLen {
		return "", ErrBadRequest
	}

	return id, nil
}
