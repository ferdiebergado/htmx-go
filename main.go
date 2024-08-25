package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ferdiebergado/htmx-go/config"
	"github.com/ferdiebergado/htmx-go/handlers"
	"github.com/ferdiebergado/htmx-go/middlewares"
)

//go:embed assets/*
var assetFs embed.FS

func main() {
	port := getPort()

	mux := http.NewServeMux()

	mux.Handle("GET /assets/", http.FileServer(http.FS(assetFs)))

	mux.HandleFunc("GET /", handlers.HandleRoot)

	mux.HandleFunc("GET /travels", handlers.HandleTravels)

	pipeline := middlewares.CreatePipeline(
		middlewares.RequestLogger,
	)

	fmt.Printf("Listening on localhost:%s...\n", port)

	log.Fatal(http.ListenAndServe(port, pipeline(mux)))
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return config.Port
	}

	return port
}
