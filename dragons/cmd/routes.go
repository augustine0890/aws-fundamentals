package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
