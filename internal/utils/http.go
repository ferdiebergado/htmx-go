package utils

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/view"
)

func IsTrustedDomain(urlStr string) bool {
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

func IsValidURL(urlStr string) bool {
	// Parse the URL to validate its format
	log.Println("Validating url...")
	_, err := url.Parse(urlStr)

	return err == nil
	// Perform additional validation if needed (e.g., check for specific schemes, protocols)

	// return true
}

func CacheBustedURL(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return "", err
	}

	// Get the file's last modified timestamp
	timestamp := fileInfo.ModTime().Unix()

	// Append the timestamp as a query parameter
	return filePath + "?v=" + time.Unix(timestamp, 0).Format("20060102150405"), nil
}

func RedirectBack(w http.ResponseWriter, r *http.Request) {
	referer := r.Header.Get("Referer")

	if !IsValidURL(referer) || !IsTrustedDomain(referer) {
		log.Println("Invalid or untrusted referer")
		w.WriteHeader(http.StatusNotFound)
		view.Render(w, r, "notfound.html", nil)
		return
	}

	http.Redirect(w, r, referer, http.StatusSeeOther)
}
