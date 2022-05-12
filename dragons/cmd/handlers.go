package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) listBuckets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	buckets, err := app.storage.GetBuckets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(buckets)
}
