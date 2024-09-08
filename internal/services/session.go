package services

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/crypto"
)

type Session struct {
	ID        string
	Data      map[string]interface{}
	ExpiresAt time.Time
}

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

type SessionKey struct{}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

func (sm *SessionManager) CreateSession(w http.ResponseWriter) (*Session, error) {
	log.Println("Creating new session...")
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sessionID, err := crypto.GenerateSecureRandomBytes()

	if err != nil {
		return nil, err
	}

	session := &Session{
		ID:        sessionID,
		Data:      make(map[string]interface{}),
		ExpiresAt: time.Now().Add(time.Duration(config.SessionDuration) * time.Minute),
	}

	sm.sessions[sessionID] = session

	http.SetCookie(w, &http.Cookie{
		Name:     config.SessionName,
		Value:    session.ID,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  session.ExpiresAt,
		Path:     "/",
	})

	return session, nil
}

func (sm *SessionManager) GetSession(r *http.Request) (*Session, error) {
	log.Println("Retrieving cookie from request...")
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	cookie, err := r.Cookie(config.SessionName)
	if err != nil {
		return nil, err
	}

	log.Println("Cookie found. Reading session info...")

	session, exists := sm.sessions[cookie.Value]

	if !exists || session.ExpiresAt.Before(time.Now()) {
		return nil, http.ErrNoCookie
	}

	log.Println("Cookie value: ", cookie.Value, " Session ID: ", session.ID)

	return session, nil
}

func (sm *SessionManager) InvalidateSession(w http.ResponseWriter, r *http.Request) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	cookie, err := r.Cookie(config.SessionName)
	if err == nil {
		delete(sm.sessions, cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     config.SessionName,
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
	}
}

func (sm *SessionManager) SetCSRFToken(session *Session) (string, error) {
	token, err := crypto.GenerateSecureRandomBytes()

	if err != nil {
		return "", err
	}

	log.Println("Setting csrf token...")
	session.Data["csrf_token"] = token
	return token, nil
}

func (sm *SessionManager) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sm.GetSession(r)

		if err != nil {

			session, err = sm.CreateSession(w)

			if err != nil {
				log.Println("unable to create a session")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			_, err := sm.SetCSRFToken(session)

			if err != nil {
				log.Println("Unable to set csrf token for the session")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		ctx := context.WithValue(r.Context(), SessionKey{}, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
