package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "Personnel Travel Monitoring System - Dashboard",
		Icon:   "dashboard",
		Header: "Overview",
	}

	utils.Render(w, r, "dashboard.html", data)
}
