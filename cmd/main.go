package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-example/api"
	"go-web-example/config"
	"go-web-example/data"
	"go-web-example/handlers"
	"go-web-example/shared"
	"log"
)

func main() {
	config.InitConfig()

	dataService, err := data.NewDatabaseService()
	if err != nil {
		log.Fatal(err)
	}

	cloudflare, err := api.NewCloudflareService()
	if err != nil {
		log.Fatal(err)
	}

	mailService, err := api.NewMailgunService()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandlers(dataService, cloudflare, mailService)

	// TODO: Remove this in production - Create users table
	_, err = dataService.DB.Exec("create table if not exists user (id text not null primary key, password text, email text, role_id text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}

	// TODO: Remove this in production - Create courses table
	_, err = dataService.DB.Exec("create table if not exists video (id text not null primary key, title text, description text, url text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}

	// TODO: Remove this in production - Create courses table
	_, err = dataService.DB.Exec("create table if not exists role (id text not null primary key, name text, permissions text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}

	// TODO: Remove this in production - Create courses table
	_, err = dataService.CreateRole("superadmin", []string{
		"ManageUsers",
		"ManageCourses",
		"AssignCourses",
		"ResetCredentials",
		"BulkUploadUsers",
		"ResetGroupCredentials",
		"ManageGroupUsers",
		"ViewCourses"})
	_, err = dataService.CreateRole("user", []string{"ViewCourses"})
	_, err = dataService.CreateRole("admin", []string{"ViewCourses", "AssignCourses", "ResetGroupCredentials", "BulkUploadUsers", "ManageGroupUsers"})

	// TODO: Remove this in production - Create courses table
	_, err = dataService.CreateUser("user@time", "passtime", "user")
	_, err = dataService.CreateUser("admin@time", "passtime", "admin")
	if err != nil {
		log.Fatal(err)
	}

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
	e.GET("/login", handler.LoginIndex, handler.RedirectIfLoggedInMiddleware)
	e.POST("/login", handler.Login, handler.RedirectIfLoggedInMiddleware)
	e.GET("/register", handler.RegisterIndex, handler.RedirectIfLoggedInMiddleware)
	e.POST("/register", handler.Register, handler.RedirectIfLoggedInMiddleware)
	e.POST("/logout", handler.Logout)

	e.GET("/courses", handler.Course)
	//e.GET("/")

	e.Logger.Fatal(e.Start("localhost:3000")) // TODO REMOVE IN PRODUCTION
}
