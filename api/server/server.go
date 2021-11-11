package server

import (
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	bootstrap(r)
	return r
}
