package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) listBuckets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	buckets, err := app.storage.GetBuckets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(buckets)
}

func (app *Application) queryBucket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty Bucket Name", http.StatusBadRequest)
		return
	}
	payload, err := app.storage.QueryBucket(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(string(payload))
}
