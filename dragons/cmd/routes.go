package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) routes() *mux.Router {
	routers := mux.NewRouter()

	s3 := routers.PathPrefix("/s3").Subrouter()
	s3.HandleFunc("/buckets", app.listBuckets)
	s3.HandleFunc("/bucket/items", app.listItems)
	s3.HandleFunc("/bucket", app.createBucket).Methods(http.MethodPost)
	s3.HandleFunc("/bucket", app.queryBucket).Methods(http.MethodGet)
	s3.HandleFunc("/bucket", app.deleteBucket).Methods(http.MethodDelete)

	return routers
}
