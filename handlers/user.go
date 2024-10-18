package handlers

import (
	"github.com/labstack/echo/v4"
	"go-web-example/views/users"
)

func (u *Handler) ManageUsers(c echo.Context) error {
	return Render(c, users.ManageUsers())
}
