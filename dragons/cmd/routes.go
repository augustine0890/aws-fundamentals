package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/s3/list", app.listBuckets)
	mux.HandleFunc("/s3/bucket", app.queryBucket)
	return mux
}
