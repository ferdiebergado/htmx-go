package handlers

import (
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/view"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	view.Render(w, r, "home.html", nil)
}
