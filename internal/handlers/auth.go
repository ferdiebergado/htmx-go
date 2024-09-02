package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "login.html", nil)
}
