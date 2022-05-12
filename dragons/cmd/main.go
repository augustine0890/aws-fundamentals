package main

import (
	"flag"
	"log"
	"net/http"

	"dragons/pkg/storage"
)

type Application struct {
	storage *storage.Storage
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	app := &Application{
		storage: storage.NewStorage(""),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
