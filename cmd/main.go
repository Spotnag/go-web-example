package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-example/api"
	"go-web-example/data"
	"go-web-example/handlers"
	"go-web-example/shared"
	"log"
	"os"
)

func main() {
	os.Setenv("CLOUDFLARE_API_TOKEN", "vJPpC2AdUZeUb8B2nzSbZePe_UZ74DAlLzriGD9f")
	os.Setenv("CLOUDFLARE_ACCOUNT_ID", "8b255e52c471b0801a25f6ecfc31ae07")

	dataService, err := data.NewDataService()
	if err != nil {
		log.Fatal(err)
	}

	apiService, err := api.NewApiService()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandlers(dataService, apiService)

	// TODO: Remove this in production - Create users table
	_, err = dataService.DB.Exec("create table if not exists user (id text not null primary key, username text, password text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}
	_, err = dataService.CreateUser("admin@time", "passtime")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Remove this in production - Create courses table
	_, err = dataService.DB.Exec("create table if not exists video (id text not null primary key, title text, description text, url text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}

	//course := []*models.Course{
	//	{
	//		ID:            "1",
	//		Title:         "Course 1",
	//		Author:        "Author 1",
	//		PublishedDate: "2021-01-01",
	//	},
	//}

	e := echo.New()

	//e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handler.CheckLoggedInMiddleware)
	//e.Use(middleware.CSRF()) // TODO FIX
	e.HTTPErrorHandler = shared.CustomHTTPErrorHandler
	e.Debug = true // TODO REMOVE IN PRODUCTION
	e.RouteNotFound("/*", shared.MissingRouteHandler)

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", handler.HandleHome)
	e.GET("/login", handler.LoginIndex)
	e.GET("/register", handler.RegisterIndex)
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)
	e.POST("/register", handler.Register)

	e.GET("/courses", handler.Course)

	e.Logger.Fatal(e.Start("localhost:3000")) // TODO REMOVE IN PRODUCTION
}
