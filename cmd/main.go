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
	"github.com/ferdiebergado/htmx-go/internal/utils"

	_ "github.com/lib/pq"
)

func main() {
	utils.LoadEnv()
	log.Println("APP_KEY: ", os.Getenv("APP_KEY"))

	port := cmp.Or(os.Getenv("APP_PORT"), config.Port)
	assetsPath := fmt.Sprintf("/%s/", config.AssetsDir)

	database := db.GetDb()
	defer database.Close()

	// opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(opt)

	// redisClient := redis.NewClient(opt)

	sessionManager := services.NewSessionManager()

	baseMiddlewares := []router.Middleware{
		middlewares.RequestLogger,
		sessionManager.SessionMiddleware,
	}

	withCsrf := append(baseMiddlewares, sessionManager.CSRFMiddleware)

	router := router.NewRouter()

	// assets
	router.Handle("GET"+assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(config.AssetsDir))))

	// root
	router.Handle("GET /", http.HandlerFunc(handlers.ShowDashboard))

	// activities
	activityHandler := handlers.ActivityHandler{Repository: models.NewActivityRepository(database)}

	router.Handle("GET /activities", http.HandlerFunc(activityHandler.ListActivities), baseMiddlewares...)
	router.Handle("GET /activities/new", http.HandlerFunc(activityHandler.ShowActivityForm), baseMiddlewares...)
	router.Handle("POST /activities", http.HandlerFunc(activityHandler.CreateActivity), withCsrf...)
	router.Handle("GET /activities/{id}", http.HandlerFunc(activityHandler.ShowActivity), baseMiddlewares...)

	router.Handle("GET /personnel", http.HandlerFunc(handlers.HandlePersonnel), baseMiddlewares...)
	router.Handle("GET /travels", http.HandlerFunc(handlers.HandleTravels), baseMiddlewares...)

	router.Handle("GET /login", http.HandlerFunc(handlers.Login), baseMiddlewares...)

	fmt.Printf("Listening on localhost:%s...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), middlewares.ErrorHandler(router)))
}
