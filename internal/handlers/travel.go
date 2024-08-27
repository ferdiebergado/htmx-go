package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func HandleTravels(w http.ResponseWriter, _ *http.Request) {

	data := &PageData{
		Title:  "PTMS - Travels",
		Icon:   "plane",
		Header: "Travels",
	}

	utils.Render(w, "travel.html", data)

}
