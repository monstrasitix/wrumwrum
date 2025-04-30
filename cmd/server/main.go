package main

import (
	"log"
	"net/http"
	"sim/internal/routing"
	"sim/internal/web"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	web.Setup()
	mux.Handle("/static/", routing.StaticHandler())

	mux.Handle("/", routing.HomeHandler{})
	mux.Handle("/login", routing.LoginHandler{})

	log.Printf("Listening on: http://localhost%s", server.Addr)
	server.ListenAndServe()
}
