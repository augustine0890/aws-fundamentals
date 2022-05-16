package main

import (
	"dragons/internal/storage"
	"flag"
	"log"
	"net/http"

	transportHttp "dragons/internal/transport/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	storageService := storage.NewStorage("")

	httpHandler := transportHttp.NewHandler(storageService)

	srv := &http.Server{
		Addr:    *addr,
		Handler: httpHandler.Router,
	}
	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
