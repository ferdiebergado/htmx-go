package utils

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/ferdiebergado/htmx-go/internal/config"
	"github.com/ferdiebergado/htmx-go/internal/services"
)

var templateCache = sync.Map{}

func logResult(templateName string, isHit bool) {
	result := "miss"

	if isHit {
		result = "hit"
	}

	log.Printf("Template cache %s: %s", result, templateName)
}

func parseTemplate(templateName string, r *http.Request) *template.Template {

	if cachedTemplate, ok := templateCache.Load(templateName); ok {
		logResult(templateName, true)
		return cachedTemplate.(*template.Template)
	}

	logResult(templateName, false)

	layoutPath := config.TemplatesDir + "/" + config.MasterTemplate
	templatePath := config.TemplatesDir + "/" + templateName

	session, ok := r.Context().Value(services.SessionKey{}).(*services.Session)

	if !ok {
		session = &services.Session{}
	}

	var getCsrf = func() string {
		return session.Data["csrf_token"].(string)
	}

	var getSession = func() *services.Session {
		return session
	}

	funcs := template.FuncMap{"csrf_token": getCsrf, "session": getSession}

	tmpl := template.Must(template.New("").Funcs(funcs).ParseFiles(layoutPath, templatePath))

	templateCache.Store(templateName, tmpl)

	return tmpl
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {

	t := parseTemplate(tmpl, r)

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, tmpl, data); err != nil {
		panic(err)
	}

	buf.WriteTo(w)
}
