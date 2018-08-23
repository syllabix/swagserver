# swagserver
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/swagserver)](https://goreportcard.com/report/github.com/syllabix/swagserver)

Middleware used to serve swagger ui with go-swagger generated api servers

## Usage
This simple project is intended to be used as middle ware that can be used to serve swagger ui using a valid Open API json specification document.
The middleware only has standard lib dependencies, so it should be fairly easy to compose in the context of other popular web frameworks.

## Setup and usage

```go get github.com/syllabix/swagserver```

Initialize an instance of swag server middle ware with the new factory function using a `swagserver.Config` struct. Register it as middleware that occurs prior to routing.

```go
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	swagserve := swagserver.New(swagserver.Config{
		URLPath:        "/v1/swagger",
		ShouldRender:   true,
		SwaggerSpecURL: "/swagger.json",
	})
	return swagserve(handler)
}
```
## Development
To remove the need to keep static web ui files in the repository, this project uses [statik](https://github.com/rakyll/statik) to embed the [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist) npm package into the package's binary. A simple build script is included in the repository that can be used to resolve [swagger-ui-dist](https://www.npmjs.com/package/swagger-ui-dist) from npm, copy and embed necessary files from the package.



