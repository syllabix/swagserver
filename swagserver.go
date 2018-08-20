//go:generate statik -src=./tmp
package swagserver

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rakyll/statik/fs"

	// This empty import is required to initialize the swagger ui data
	_ "github.com/syllabix/swagserver/statik"
)

// Middleware is a simple alias for a commonmiddle ware function signature, func(next http.Handler) http.Handler
type Middleware func(next http.Handler) http.Handler

// Config is an object intended to be used to configure
// an instance of a swag server middle ware
type Config struct {

	// URLPath is the path at which swagger ui will be served at
	URLPath string

	// ShouldRender is used to signal to the middleware whether or not to serve
	// swagger ui (Ex. expose swagger ui in non production environments, while not rendering in production)
	ShouldRender bool

	// SwaggerSpecURL is the url swagger ui will try to resolve a valid specifcation file from
	SwaggerSpecURL string
}

var (
	shouldServeSwagger func(url *url.URL) bool

	handler *swaghandler

	defaultConfig = Config{
		URLPath:        "/swagger",
		ShouldRender:   true,
		SwaggerSpecURL: "/swagger.json",
	}
)

func handleError(err error) {
	if err != nil {
		log.Fatal("An error occurred while initializing swagserver middlware:", err)
	}
}

func init() {
	statikFs, err := fs.New()
	handleError(err)

	swaggerui, err := statikFs.Open("/index.html")
	handleError(err)
	defer swaggerui.Close()

	data, err := ioutil.ReadAll(swaggerui)
	handleError(err)

	fileserver := http.FileServer(statikFs)
	tmpl, err := template.New("swaggerui").Parse(string(data))
	handleError(err)

	handler = newSwagHandler(fileserver, tmpl)
}

func setConfig(configurations []Config) Config {
	switch len(configurations) {
	case 0:
		return defaultConfig
	case 1:
		return configurations[0]
	default:
		log.Println(`Multiple configuration objects received,
		taking the first setting provided and ignoring the rest`)
		return configurations[0]
	}
}

// New is the factory constructor for intializing a SwagServer middleware
// It takes an optional SwagServerConfig object as its single argument.
// If used without providing a config object, it will intialize the middleware
// using default settings:
// 	UrlPath:      "/swagger/"
// 	ShouldRender: true
// 	SpecUrl:      "/swagger.json"
func New(config ...Config) Middleware {
	cfg := setConfig(config)

	shouldServeSwagger = func(url *url.URL) bool {
		return cfg.ShouldRender && strings.HasPrefix(url.Path, cfg.URLPath)
	}

	handler.specURL = cfg.SwaggerSpecURL

	return func(next http.Handler) http.Handler {
		return serveSwagger(next, cfg.SwaggerSpecURL)
	}
}

func serveSwagger(next http.Handler, swaggerURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if shouldServeSwagger(r.URL) {
			handler.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
