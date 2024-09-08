package utils

import (
	"log"
	"net/http"
	"net/url"

	"github.com/ferdiebergado/htmx-go/internal/config"
)

func RedirectBack(w http.ResponseWriter, r *http.Request) {
	referer := r.Header.Get("Referer")

	if !isValidURL(referer) || !isTrustedDomain(referer) {
		log.Println("Invalid or untrusted referer")
		w.WriteHeader(http.StatusNotFound)
		Render(w, r, "notfound.html", nil)
		return
	}

	http.Redirect(w, r, referer, http.StatusSeeOther)
}

func isTrustedDomain(urlStr string) bool {
	// Parse the URL to extract the domain
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	log.Println("PARSEDURL:", parsedURL, parsedURL.Host)

	// Define a list of trusted domains
	trustedDomains := config.TrustedDomains

	// Check if the domain matches any of the trusted domains
	for _, trustedDomain := range trustedDomains {
		if parsedURL.Host == trustedDomain {
			return true
		}
	}

	return false
}

func isValidURL(urlStr string) bool {
	// Parse the URL to validate its format
	log.Println("Validating url...")
	_, err := url.Parse(urlStr)

	return err == nil
	// Perform additional validation if needed (e.g., check for specific schemes, protocols)

	// return true
}
