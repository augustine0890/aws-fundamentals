package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/s3/list", app.listBuckets)
	mux.HandleFunc("/s3/bucket", app.queryBucket)
	mux.HandleFunc("/s3/create", app.createBucket)

	return mux
}
