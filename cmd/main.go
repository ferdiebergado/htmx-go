package main

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/db"
	"github.com/ferdiebergado/htmx-go/internal/handlers"
	"github.com/ferdiebergado/htmx-go/internal/middlewares"
	"github.com/ferdiebergado/htmx-go/internal/models"
	"github.com/ferdiebergado/htmx-go/internal/router"
	"github.com/ferdiebergado/htmx-go/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	// port
	port := cmp.Or(os.Getenv("APP_PORT"), config.Port)

	// database
	database := db.GetDb()
	defer database.Close()

	// session manager
	sessionManager := services.NewSessionManager()

	// app
	app := router.NewRouter()

	// base middlewares
	app.RegisterMiddlewares(middlewares.RequestLogger, sessionManager.SessionMiddleware, middlewares.CSRFMiddleware)

	// assets
	app.Handle(router.GET, config.AssetsPath, http.StripPrefix(config.AssetsPath, http.FileServer(http.Dir(config.AssetsDir))))

	// root
	app.Handle(router.GET, "/", http.HandlerFunc(handlers.HomeHandler))

	// dashboard
	app.Handle(router.GET, "/dashboard", http.HandlerFunc(handlers.ShowDashboard))

	// activities
	activityHandler := handlers.ActivityHandler{Repository: models.NewActivityRepository(database)}

	app.Handle(router.GET, "/activities", http.HandlerFunc(activityHandler.ListActivities))
	app.Handle(router.GET, "/activities/new", http.HandlerFunc(activityHandler.ShowActivityForm))
	app.Handle(router.POST, "/activities", http.HandlerFunc(activityHandler.CreateActivity))
	app.Handle(router.GET, "/activities/{id}", http.HandlerFunc(activityHandler.ShowActivity))
	app.Handle(router.GET, "/activities/{id}/edit", http.HandlerFunc(activityHandler.ShowActivityEditForm))
	app.Handle(router.POST, "/activities/{id}", http.HandlerFunc(activityHandler.UpdateActivity))
	app.Handle(router.POST, "/activities/{id}/delete", http.HandlerFunc(activityHandler.DestroyActivity))

	// personnel
	app.Handle(router.GET, "/personnel", http.HandlerFunc(handlers.HandlePersonnel))

	// travels
	app.Handle(router.GET, "/travels", http.HandlerFunc(handlers.HandleTravels))

	// auth
	app.Handle(router.GET, "/login", http.HandlerFunc(handlers.Login))

	// server
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      middlewares.ErrorHandler(app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on localhost:%s...\n", port)
	log.Fatal(srv.ListenAndServe())
}
