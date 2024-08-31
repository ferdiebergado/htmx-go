package utils

import (
	"bytes"
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

func parseTemplate(templateName string) *template.Template {

	if cachedTemplate, ok := templateCache.Load(templateName); ok {
		logResult(templateName, true)
		return cachedTemplate.(*template.Template)
	}

	logResult(templateName, false)

	layoutPath := config.TemplatesDir + "/" + config.MasterTemplate
	templatePath := config.TemplatesDir + "/" + templateName

	tmpl := template.Must(template.ParseFiles(layoutPath, templatePath))

	templateCache.Store(templateName, tmpl)

	return tmpl
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	t := parseTemplate(tmpl)

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, tmpl, data); err != nil {
		panic(err)
	}

	buf.WriteTo(w)
}
