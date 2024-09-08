package config

var TrustedDomains = []string{"localhost:8080", "yourdomain.com"}

var Database = map[string]interface{}{
	"Driver": "postgres",
}

const (
	Port            string = "8888"
	AssetsPath      string = "/static/"
	AssetsDir       string = "public"
	TemplatesDir    string = "templates"
	MasterTemplate  string = "layout.html"
	SessionName     string = "sid"
	SessionDuration int    = 30
)
