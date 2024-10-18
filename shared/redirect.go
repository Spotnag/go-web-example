package shared

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HXRedirect(c echo.Context, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", url)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}
	return c.Redirect(http.StatusSeeOther, url)
}

func HXLocation(c echo.Context, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Location", url)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
