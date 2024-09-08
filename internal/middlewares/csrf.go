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
			log.Println("CSRFtoken from form:", csrfToken)

			if !ValidateCSRFToken(session, csrfToken) {
				log.Println("Invalid CSRF token")
				w.WriteHeader(http.StatusForbidden)
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
		log.Println("csrftoken from session:", storedToken)
		return storedToken == token
	}
	return false
}
