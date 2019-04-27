package option

import (
	"github.com/syllabix/swagserver/config"
	"github.com/syllabix/swagserver/theme"
)

// Path is an option used to set the path the middleware should
// serve swagger ui at
func Path(path string) config.Option {
	return func(s *config.Settings) {
		s.URLPath = path
	}
}

// Disable the middleware from serving swagger ui
func Disable() config.Option {
	return func(s *config.Settings) {
		s.ShouldRender = false
	}
}

// SwaggerSpecURL configures the middleware to resolve the swagger spec file from
// the provided url
func SwaggerSpecURL(url string) config.Option {
	return func(s *config.Settings) {
		s.SwaggerSpecURL = url
	}
}

// Theme is option that used to configure the styling of the served swagger ui
func Theme(name theme.Name) config.Option {
	return func(s *config.Settings) {
		s.Theme = name
	}
}
