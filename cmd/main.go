package main

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"

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
	log.Println("APP_KEY: ", os.Getenv("APP_KEY"))

	port := cmp.Or(os.Getenv("APP_PORT"), config.Port)

	database := db.GetDb()
	defer database.Close()

	sessionManager := services.NewSessionManager()

	pipeline := middlewares.CreatePipeline(
		middlewares.ErrorHandler,
	)

	app := router.NewRouter()

	app.RegisterMiddlewares(middlewares.RequestLogger, sessionManager.SessionMiddleware, middlewares.CSRFMiddleware)

	// assets
	app.Handle(router.GET, config.AssetsPath, http.StripPrefix(config.AssetsPath, http.FileServer(http.Dir(config.AssetsDir))))

	// root
	app.Handle(router.GET, "/", http.HandlerFunc(handlers.HomeHandler))
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

	app.Handle(router.GET, "/personnel", http.HandlerFunc(handlers.HandlePersonnel))
	app.Handle(router.GET, "/travels", http.HandlerFunc(handlers.HandleTravels))

	app.Handle(router.GET, "/login", http.HandlerFunc(handlers.Login))

	fmt.Printf("Listening on localhost:%s...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), pipeline(app)))
}
