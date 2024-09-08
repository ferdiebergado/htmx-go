package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/view"
)

func HandleTravels(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "PTMS - Travels",
		Icon:   "plane",
		Header: "Travels",
	}

	view.Render(w, r, "travel.html", data)

}
