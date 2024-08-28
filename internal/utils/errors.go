package utils

import (
	"fmt"
	"log"
	"net/http"
)

func LogError(err error) {
	log.Println(fmt.Errorf("error: %v", err))
}

func HandleHTTPError(w http.ResponseWriter, err error) {
	LogError(err)
	w.WriteHeader(http.StatusInternalServerError)
	Render(w, "error.html", nil)
}
