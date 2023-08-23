package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/marioteik/website/controllers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.Heartbeat("/ping"))
	//mux.Use(NoSurf)
	//mux.Use(SessionLoad)

	staticFileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", staticFileServer))

	mux.Mount("/", controllers.LandingPage{}.Routes())

	return mux
}
