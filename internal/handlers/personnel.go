package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func HandlePersonnel(w http.ResponseWriter, r *http.Request) {

	data := &PageData{
		Title:  "PTMS - Personnel",
		Icon:   "users",
		Header: "Personnel",
	}

	utils.Render(w, r, "personnel.html", data)

}
