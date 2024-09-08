package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ferdiebergado/htmx-go/internal/utils"
	"github.com/ferdiebergado/htmx-go/internal/view"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				// Handle generic panics
				log.Println("Unexpected error:", err)
				log.Println(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				view.Render(w, r, "error.html", nil)
			}
		}()

		cw := &customWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(cw, r)

		switch cw.statusCode {
		case http.StatusNotFound:
			log.Println("R.URL:", r.URL, ".PATH:", r.URL.Path)
			view.Render(w, r, "notfound.html", nil)
		case http.StatusInternalServerError:
			view.Render(w, r, "error.html", nil)
		case http.StatusBadRequest:
			// TODO:
			utils.RedirectBack(w, r)
		case http.StatusForbidden:
			view.Render(w, r, "forbidden.html", nil)
		}
	})
}
