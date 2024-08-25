package templates

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"sync"

	"github.com/ferdiebergado/htmx-go/config"
)

//go:embed *.html
var TemplateFS embed.FS

var templateCache = sync.Map{}

func parseTemplate(templateName string) (*template.Template, error) {
	if cachedTemplate, ok := templateCache.Load(templateName); ok {
		return cachedTemplate.(*template.Template), nil
	}

	tmpl, err := template.ParseFS(TemplateFS, config.MasterTemplate, templateName)

	if err != nil {
		return nil, err
	}

	templateCache.Store(templateName, tmpl)

	return tmpl, nil
}

func Render(w http.ResponseWriter, tmpl string, data any) {
	t, err := parseTemplate(tmpl)

	if err != nil {
		http.Error(w, "Unable to parse html files", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, tmpl, data); err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
