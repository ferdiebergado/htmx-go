package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func HandleActivities(w http.ResponseWriter, _ *http.Request) {

	data := &PageData{
		Title:  "PTMS - Activities",
		Icon:   "calendar-o",
		Header: "Activities",
	}

	utils.Render(w, "activities.html", data)

}
