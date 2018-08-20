package swagserver

import (
	"html/template"
	"net/http"
	"path"
	"regexp"
)

type swaghandler struct {
	fileserver      http.Handler
	tmpl            *template.Template
	staticFileRegex *regexp.Regexp
	specURL         string
}

func newSwagHandler(fileserver http.Handler, tmpl *template.Template) *swaghandler {
	return &swaghandler{
		fileserver:      fileserver,
		staticFileRegex: regexp.MustCompile(`\.(?:j|cs)s`),
		tmpl:            tmpl,
	}
}

func (s *swaghandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.staticFileRegex.MatchString(r.URL.Path) {
		dir, _ := path.Split(r.URL.Path)
		http.StripPrefix(dir, s.fileserver).ServeHTTP(w, r)
	} else {
		s.tmpl.Execute(w, s.specURL)
	}
}
