package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/view"
)

func Login(w http.ResponseWriter, r *http.Request) {
	view.Render(w, r, "login.html", nil)
}
