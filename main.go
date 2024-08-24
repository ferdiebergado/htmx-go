package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferdiebergado/htmx-go/handlers"
	"github.com/ferdiebergado/htmx-go/middlewares"
)

func main() {

	port := ":8888"

	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.HandleRoot)

	fmt.Printf("Listening on localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(port, middlewares.RequestLogger(router)))
}
