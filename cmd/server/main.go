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
	mux.Handle("/table", routing.TableHandler{})
	mux.Handle("/charts", routing.ChartsHandler{})

	log.Printf("Listening on: http://localhost%s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
