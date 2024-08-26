package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "Personnel Travel Monitoring System - Dashboard",
		Icon:   "fa-dashboard",
		Header: "Overview",
	}

	utils.Render(w, "dashboard.html", data)
}
