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

func parseTemplate(templateName string) (*template.Template, error) {
	const cacheTxt string = "Template cache"
	const templatePathFormat string = "%s/%s"

	if cachedTemplate, ok := templateCache.Load(templateName); ok {
		log.Printf("%s hit: %s", cacheTxt, templateName)
		return cachedTemplate.(*template.Template), nil
	}

	log.Printf("%s miss: %s", cacheTxt, templateName)

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
		log.Println(fmt.Errorf("%w", err))
		http.Error(w, "Unable to parse html files", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, tmpl, data); err != nil {
		log.Println(fmt.Errorf("%w", err))
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
