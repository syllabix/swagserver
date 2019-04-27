package swagserver

import (
	"html/template"
	"net/http"
	"path"
	"strings"

	"github.com/syllabix/swagserver/config"
	"github.com/syllabix/swagserver/theme"
)

type tmpldata struct {
	SpecURL  string
	Theme    theme.Name
	BasePath string
}

type handler struct {
	fileserver http.Handler
	tmpl       *template.Template
	specURL    string
	basePath   string
	theme      theme.Name
}

func isStaticFile(path string) bool {
	return strings.HasSuffix(path, ".js") ||
		strings.HasSuffix(path, ".css") ||
		strings.HasSuffix(path, ".png")
}

func newhandler(settings config.Settings, fileserver http.Handler, tmpl *template.Template) *handler {
	return &handler{
		fileserver: fileserver,
		tmpl:       tmpl,
		specURL:    settings.SwaggerSpecURL,
		basePath:   strings.TrimRight(settings.URLPath, "/"),
		theme:      settings.Theme,
	}
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isStaticFile(r.URL.Path) {
		dir, _ := path.Split(r.URL.Path)
		http.StripPrefix(dir, s.fileserver).ServeHTTP(w, r)
	} else {
		s.tmpl.Execute(w, tmpldata{
			SpecURL:  s.specURL,
			Theme:    s.theme,
			BasePath: s.basePath,
		})
	}
}
