package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

