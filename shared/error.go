package shared

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		if err := c.String(he.Code, he.Message.(string)); err != nil {
			c.Logger().Error(err)
		}
		return
	}
	c.Logger().Error(err) // log the original error
	if err = c.String(http.StatusInternalServerError, "internal server error"); err != nil {
		c.Logger().Error(err)
	}
}

func MissingRouteHandler(c echo.Context) error {
	return c.HTML(http.StatusNotFound, "404 Page does not exist...")
}
