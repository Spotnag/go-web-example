package handlers

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go-web-example/shared"
	"go-web-example/views/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func (u *Handler) LoginIndex(c echo.Context) error {
	return Render(c, auth.Login())
}

func (u *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check if the password is correct
	targetUser, err := u.DataService.GetUser(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Username")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Password")
	}

	// Set user as authenticated
	session, _ := store.Get(c.Request(), "session")
	session.Values["loggedIn"] = true
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   48 * 60 * 60, // 48 hours
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode, // TODO what are the differences here?
	}
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return shared.HXRedirect(c, "/")

}

func (u *Handler) Logout(c echo.Context) error {
	session, _ := store.Get(c.Request(), "session")

	// Revoke users authentication
	session.Values["loggedIn"] = false
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return shared.HXRedirect(c, "/")
}

func (u *Handler) CheckLoggedInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isLoggedIn := session.Values["loggedIn"]
		if isLoggedIn == nil {
			isLoggedIn = false
		}
		c.SetRequest(c.Request().WithContext(context.WithValue(
			c.Request().Context(),
			"isLoggedIn",
			isLoggedIn)))
		return next(c)
	}
}
