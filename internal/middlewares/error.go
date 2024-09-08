package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				// Handle generic panics
				log.Println("Unexpected error:", err)
				log.Println(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				utils.Render(w, r, "error.html", nil)
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
			if !strings.HasPrefix(r.URL.Path, config.AssetsPath) {
				utils.Render(w, r, "notfound.html", nil)
			}
		case http.StatusInternalServerError:
			utils.Render(w, r, "error.html", nil)
		case http.StatusBadRequest:
			utils.RedirectBack(w, r)
		case http.StatusForbidden:
			utils.Render(w, r, "forbidden.html", nil)
		}
	})
}
