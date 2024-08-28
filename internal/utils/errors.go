package utils

import (
	"fmt"
	"log"
	"net/http"
)

func LogError(desc string, err error) {
	details := fmt.Errorf("error: %v", err)
	log.Printf("ERROR: %s %s", desc, details)
}

func HandleHTTPError(w http.ResponseWriter, desc string, err error) {
	LogError(desc, err)
	w.WriteHeader(http.StatusInternalServerError)
	Render(w, "error.html", nil)
}
