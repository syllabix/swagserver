package config

import "github.com/syllabix/swagserver/theme"

// An Option is used to configure swag server settings
type Option func(s *Settings)

// Settings is an object intended to be used to configure
// an instance of a swag server middle ware
type Settings struct {

	// URLPath is the path at which swagger ui will be served at
	URLPath string

	// ShouldRender is used to signal to the middleware whether or not to serve
	// swagger ui (Ex. expose swagger ui in non production environments, while not rendering in production)
	ShouldRender bool

	// SwaggerSpecURL is the url swagger ui will try to resolve a valid specifcation file from
	SwaggerSpecURL string

	// Theme is the theme that will be applied to rendered swagger ui
	Theme theme.Name
}

// DefaultSettings used to configure a swag serve middleware
var DefaultSettings = Settings{
	URLPath:        "/swagger/",
	ShouldRender:   true,
	SwaggerSpecURL: "/swagger.json",
	Theme:          theme.Standard,
}
