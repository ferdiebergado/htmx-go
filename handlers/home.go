package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/templates"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title  string
		Icon   string
		Header string
	}{
		Title:  "Personnel Travel Monitoring System - Dashboard",
		Icon:   "fa-dashboard",
		Header: "Overview",
	}

	templates.Render(w, "dashboard.html", data)
}
