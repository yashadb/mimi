package mimi

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/zulubit/mimi/pkg/router"
)

type MimiExtend struct {
	ExtraRoutes       []router.ExtraRoute
	TemplateFunctions map[string]interface{}
	// TODO: Implement logging
}

func (m MimiExtend) Start() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, router.SetupRouter(m.ExtraRoutes))
	err := http.ListenAndServe(":8080", loggedRouter)
	panic(err)
}
