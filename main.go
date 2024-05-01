package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"go-web-example/handlers"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	logger := httplog.NewLogger("logger", httplog.Options{
		//JSON:             true, // TODO enable in prod
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})
	r.Mount("/courses", CourseRoutes())
	srv := http.Server{
		Addr:    "localhost:3000",
		Handler: r,

		//These make sure the server does not hang
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,

		// this prevents Slowloris attack
		IdleTimeout: time.Second * 60,
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}

func CourseRoutes() chi.Router {
	r := chi.NewRouter()
	courseHandler := handlers.CourseHandler{}
	r.Get("/", courseHandler.ListCourses)
	r.Post("/", courseHandler.CreateCourse)
	r.Get("/{ID}", courseHandler.GetCourses)
	r.Put("/{ID}", courseHandler.UpdateCourse)
	r.Delete("/{ID}", courseHandler.DeleteCourse)
	return r
}
