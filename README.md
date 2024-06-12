# swagserver

[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/swagserver)](https://goreportcard.com/report/github.com/syllabix/swagserver)

Swagserver is a simple project designed to serve your swagger ui using a valid specification file. It can be added as a middleware or registered as a dedicated http handler. It was intially built to be using with [go swagger](https://github.com/go-swagger/go-swagger) generated servers - but should be pluggable with with most routing frameworks.

## Swagger UI Versions

Curious what swagger ui version we are on?

The corresponding swagger ui version is included in the swagserver [release notes](https://github.com/syllabix/swagserver/releases).

## Setup and usage

```
go get github.com/syllabix/swagserver@latest
```

### As a handler with a go swagger generated API

```go
func main() {

	// initialize a new swag server handler
	swagger := swagserver.NewHandler(
		option.SwaggerSpecURL("/swagger.json"),
		option.Theme(theme.Muted),
		option.Path("/docs/"),
	)

	// initialize your go swagger server
	spec, err := loads.Analyzed(rest.FlatSwaggerJSON, "")
	if err != nil {
		log.Fatal("nope... not this time", err)
	}
	api := operation.NewAwesomeAPI(spec)
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()


	mux := http.NewServeMux()

	// register swagger as a hendler
	mux.Handle("/docs/", swagger)

	// otherwise serve the api
	mux.Handle("/", api.Serve(nil))

	server := http.Server{
		Addr:    ":9091",
		Handler: mux,
	}

	fmt.Println("starting server")
	server.ListenAndServe()
}
```

### As a middleware

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

A simple build script is included in the repository that can be used to resolve [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist) from npm, copy and embed necessary files from the package. You will need to be running on a mac (macOS style `sed` is used) and have npm and statik installed in order for it to run.
