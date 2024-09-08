package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/view"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "Personnel Travel Monitoring System - Dashboard",
		Icon:   "dashboard",
		Header: "Overview",
	}

	view.Render(w, r, "dashboard.html", data)
}
