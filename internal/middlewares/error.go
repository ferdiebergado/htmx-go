package middlewares

import (
	"log"
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n", err)
				// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				utils.Render(w, "error.html", nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
