package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/handlers"
	"github.com/ferdiebergado/htmx-go/internal/middlewares"
)

func main() {
	port := getPort()
	// assetsPath := fmt.Sprintf("/%s/", config.AssetsDir)

	mux := http.NewServeMux()

	// mux.Handle("GET "+assetsPath, http.StripPrefix(assetsPath, http.FileServer(http.Dir(config.AssetsDir))))
	mux.HandleFunc("GET /activities", handlers.HandleActivities)
	mux.HandleFunc("GET /activities/create", handlers.CreateActivity)
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

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return config.Port
	}

	return port
}
