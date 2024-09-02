package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ferdiebergado/htmx-go/internal/utils"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Errorf("error: %v", err))
				log.Println(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				utils.Render(w, "error.html", nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
