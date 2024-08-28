package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/ferdiebergado/htmx-go/internal/config"
)

var templateCache = sync.Map{}

func logResult(templateName string, isHit bool) {
	result := "miss"

	if isHit {
		result = "hit"
	}

	log.Printf("Template cache %s: %s", result, templateName)
}

func parseTemplate(templateName string) (*template.Template, error) {

	const templatePathFormat string = "%s/%s"

	if cachedTemplate, ok := templateCache.Load(templateName); ok {
		logResult(templateName, true)
		return cachedTemplate.(*template.Template), nil
	}

	logResult(templateName, false)

	layoutPath := fmt.Sprintf(templatePathFormat, config.TemplatesDir, config.MasterTemplate)
	templatePath := fmt.Sprintf(templatePathFormat, config.TemplatesDir, templateName)

	tmpl, err := template.ParseFiles(layoutPath, templatePath)

	if err != nil {
		return nil, err
	}

	templateCache.Store(templateName, tmpl)

	return tmpl, nil
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := parseTemplate(tmpl)

	if err != nil {
		HandleHTTPError(w, "unable to parse files", err)
		return
	}

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, tmpl, data); err != nil {
		HandleHTTPError(w, "unable to execute template", err)
		return
	}

	buf.WriteTo(w)
}
