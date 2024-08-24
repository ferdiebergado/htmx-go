package handlers

import (
	"net/http"
	"text/template"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/dashboard.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Personnel Travel Monitoring System",
	}

	tmpl.Execute(w, data)
}
