package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/zulubit/mimi/pkg/router"
)

func main() {
	// Set up the router
	r := router.SetupRouter()

	// Logging middleware
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	fmt.Println("Mimi is alive.")

	// Start the server
	err := http.ListenAndServe(":8080", loggedRouter)
	panic(err)
}

// TODO:
// 1. figure out how to merge in seo - build seo struct, make sure config files have seo fields, reserve fields in page-config, in the template, check what fields exist and conditionally merge them in. Prefer the page config over the global one.
// 2. Add core package (Core should hold a Mimi struct, that serves as a facade for routes, template functions and logger)
// 4. Implement a post/page loop template function. Fun.
// 5. Provide a very basic admin dash.
