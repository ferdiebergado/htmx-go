package middlewares

import (
	"log"
	"net/http"

	"github.com/ferdiebergado/htmx-go/internal/services"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(services.SessionKey{}).(*services.Session)

		if r.Method == http.MethodPost || r.Method == http.MethodDelete {
			csrfToken := r.FormValue("csrf_token")

			if ValidateCSRFToken(session, csrfToken) {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}

			log.Println("csrf token valid.")
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateCSRFToken(session *services.Session, token string) bool {
	log.Println("Validating csrf token...")
	if storedToken, ok := session.Data["csrf_token"]; ok {
		return storedToken == token
	}
	return false
}
