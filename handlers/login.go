package handlers

import (
	"github.com/labstack/echo/v4"
	"go-web-example/views/auth"
)

func HandleLoginIndex(c echo.Context) error {
	return Render(c, auth.Login())
}
