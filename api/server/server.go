package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(handlers.RecoveryHandler())

	bootstrap(r)
	return r
}
