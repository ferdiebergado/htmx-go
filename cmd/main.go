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

	mux := http.NewServeMux()

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("GET /", handlers.ShowDashboard)
	mux.HandleFunc("GET /travels", handlers.HandleTravels)

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
