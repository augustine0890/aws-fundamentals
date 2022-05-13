package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) routes() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/s3/buckets", app.listBuckets)
	mux.HandleFunc("/s3/bucket/items", app.listItems)
	mux.HandleFunc("/s3/bucket", app.createBucket).Methods(http.MethodPost)
	mux.HandleFunc("/s3/bucket", app.queryBucket).Methods(http.MethodGet)
	mux.HandleFunc("/s3/bucket", app.deleteBucket).Methods(http.MethodDelete)

	return mux
}
