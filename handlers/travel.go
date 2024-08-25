package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/templates"
)

func HandleTravels(w http.ResponseWriter, _ *http.Request) {

	data := struct {
		Title  string
		Icon   string
		Header string
	}{
		Title:  "Personnel Travel Monitoring System - Travels",
		Icon:   "fa-plane",
		Header: "Travels",
	}

	templates.Render(w, "travel.html", data)

}
