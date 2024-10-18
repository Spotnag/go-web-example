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

func (u *Handler) RegisterIndex(c echo.Context) error {
	return Render(c, auth.Register())
}

func (u *Handler) Register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Check if the user already exists
	_, err := u.db.GetUser(email)
	if err == nil {
		return echo.NewHTTPError(http.StatusConflict, "User already exists")
	}

	// Create the user
	_, err = u.db.CreateUser(email, password, "user", "default")
	if err != nil {
		return err
	}

	return shared.HXRedirect(c, "/")
}

func (u *Handler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Check if the password is correct
	targetUser, err := u.db.GetUser(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Email")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(targetUser.PasswordHash), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Password")
	}

	// Set user as authenticated
	session, _ := store.Get(c.Request(), "session")
	session.Values["loggedIn"] = true
	session.Values["role"] = targetUser.Role.Name
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
	session.Options.MaxAge = -1
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return shared.HXRedirect(c, "/")
}

func (u *Handler) CheckLoggedInAndRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isLoggedIn, exists := session.Values["loggedIn"]
		if !exists {
			isLoggedIn = false
		}
		role, exists := session.Values["role"]
		if !exists {
			role = "user"
		}

		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, "isLoggedIn", isLoggedIn)
		ctx = context.WithValue(ctx, "role", role)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func (u *Handler) RedirectIfLoggedInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isLoggedIn, ok := session.Values["loggedIn"].(bool)
		if ok && isLoggedIn {
			// User is logged in, redirect them to the home page or any other page
			return shared.HXRedirect(c, "/")
		}

		// Continue with the next handler if not logged in
		return next(c)
	}
}

func (u *Handler) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isLoggedIn, ok := session.Values["loggedIn"].(bool)
		if !ok || !isLoggedIn {
			return shared.HXRedirect(c, "/login")
		}

		return next(c)
	}
}

func (u *Handler) RequireAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		role, ok := session.Values["role"].(string)
		if !ok || role != "admin" {
			return shared.MissingRouteHandler(c)
		}

		return next(c)
	}
}
