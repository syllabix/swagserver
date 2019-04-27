# swagserver
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/swagserver)](https://goreportcard.com/report/github.com/syllabix/swagserver)

A simple middleware that serves swagger ui using a valid Open API json specification document.

The middleware func is typed using the common convention `func(next http.Handler) http.Handler` - which should make it pluggable into most Go web frameworks.

## Setup and usage

```go get github.com/syllabix/swagserver```

Initialize an instance of swag server middleware with various options.

```go
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	swagserve := swagserver.New(
		option.Path("/internal/apidocs"),
		option.SwaggerSpecURL("/swagger.json"),
		option.Theme(theme.Material),
	)
	return swagserve(handler)
}
```

# Options

`Path` - the the url path swagger ui should be served at

`Disable` - swagger ui will not be rendered

`SwaggerSpecURL` - the url where the swagger json specification is served from

`Theme` - the theme the ui should by displayed in

# Themes

Themes, with exclusion of the default, are all sourced from [swagger-ui-themes](https://github.com/ostranme/swagger-ui-themes).


## Development
To remove the need to keep static web ui files in the repository, this project uses [statik](https://github.com/rakyll/statik) to embed the [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist) npm package into the project source.

A simple build script is included in the repository that can be used to resolve [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist) from npm, copy and embed necessary files from the package. You will need to be running on a mac (macOS style `sed` is used) and npm installed in order for it to run.



