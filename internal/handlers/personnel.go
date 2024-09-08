package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/view"
)

func HandlePersonnel(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "PTMS - Personnel",
		Icon:   "users",
		Header: "Personnel",
	}

	view.Render(w, r, "personnel.html", data)

}
