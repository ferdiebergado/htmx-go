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

	_ "github.com/lib/pq"
)

func main() {

	port := cmp.Or(os.Getenv("PORT"), config.Port)

	database := db.GetDb()

	defer database.Close()

	mux := http.NewServeMux()

	activityHandler := &handlers.ActivityHandler{DB: database}

	mux.HandleFunc("GET /activities", activityHandler.ListActivities)
	mux.HandleFunc("GET /activities/new", activityHandler.ShowActivityForm)
	mux.HandleFunc("POST /activities", activityHandler.CreateActivity)
	mux.HandleFunc("GET /personnel", handlers.HandlePersonnel)
	mux.HandleFunc("GET /travels", handlers.HandleTravels)
	mux.HandleFunc("GET /", handlers.ShowDashboard)

	pipeline := middlewares.CreatePipeline(
		middlewares.RequestLogger,
		middlewares.ErrorHandler,
	)

	fmt.Printf("Listening on localhost:%s...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), pipeline(mux)))
}
