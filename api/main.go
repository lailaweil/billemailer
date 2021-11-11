package main

import (
	"github.com/lailaweil/billemailer/api/server"
	"log"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:              "0.0.0.0:8080",
		Handler:           server.New(),
		WriteTimeout: time.Millisecond * 200,
		ReadTimeout:  time.Millisecond * 200,
		IdleTimeout:  time.Millisecond * 600,
	}

	log.Fatal(srv.ListenAndServe())
}