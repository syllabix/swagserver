package main

import (
	"fmt"
	"net/http"

	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
)

func main() {

	// initialize a new swag server handler
	swagger := swagserver.NewHandler(
		option.SwaggerSpecURL("https://petstore.swagger.io/v2/swagger.json"),
		// option.Theme(theme.Muted),
		option.Path("/docs/"),
	)

	mux := http.NewServeMux()

	// register swagger as a hendler
	mux.Handle("/docs/", swagger)

	server := http.Server{
		Addr:    ":9091",
		Handler: mux,
	}

	fmt.Println("starting server")
	server.ListenAndServe()
}
