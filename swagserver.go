//go:generate statik -src=./tmp
package swagserver

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rakyll/statik/fs"
	"github.com/syllabix/swagserver/config"

	// This empty import is required to initialize the embedded swagger files
	_ "github.com/syllabix/swagserver/statik"
)

// Middleware is a simple alias for a commonmiddle ware function signature, func(next http.Handler) http.Handler
type Middleware func(next http.Handler) http.Handler

const (
	fatalMsg = `
		An error occurred while initializing swagserver middlware.
		This is most likely a problem with the swagserver package itself.
		Please file an issue on github @ github.com/syllabix/swagserver`
)

func setup() (fileserver http.Handler, tmpl *template.Template, err error) {
	statikFs, err := fs.New()
	if err != nil {
		return nil, nil, err
	}

	swaggerui, err := statikFs.Open("/index.html")
	if err != nil {
		return nil, nil, err
	}
	defer swaggerui.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, swaggerui)
	if err != nil {
		return nil, nil, err
	}

	fileserver = http.FileServer(statikFs)
	tmpl, err = template.New("swaggerui").Parse(builder.String())
	return
}

// New is the factory constructor for intializing a middleware
// using config.Option funcs to override defaults settings
func New(options ...config.Option) Middleware {
	handler, settings := handlerSettings(options...)

	shouldServe := func(url *url.URL) bool {
		return settings.ShouldRender &&
			strings.HasPrefix(url.Path, settings.URLPath)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldServe(r.URL) {
				next.ServeHTTP(w, r)
			} else {
				handler.ServeHTTP(w, r)
			}
		})
	}
}

// NewHandler returns an http.Handler that serves the swagger UI. This
// constructor ignores the `Path` option
func NewHandler(options ...config.Option) http.Handler {
	handler, _ := handlerSettings(options...)
	return handler
}

func handlerSettings(options ...config.Option) (http.Handler, config.Settings) {
	fileserver, template, err := setup()
	if err != nil {
		log.Fatal(fatalMsg, err)
	}

	settings := settings(options...)

	return newhandler(settings, fileserver, template), settings
}

func settings(options ...config.Option) config.Settings {
	settings := config.DefaultSettings
	for _, opt := range options {
		opt(&settings)
	}

	return settings
}
