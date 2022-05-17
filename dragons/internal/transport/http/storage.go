package http

import (
	"dragons/internal/storage"
	"encoding/json"
	"net/http"
)

func (h *Handler) listBuckets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	buckets, err := h.Storage.GetBuckets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(buckets)
}

func (h *Handler) queryBucket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty Bucket Name", http.StatusBadRequest)
		return
	}
	payload, err := h.Storage.QueryBucket(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(payload)
}

func (h *Handler) createBucket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty Bucket Name", http.StatusBadRequest)
		return
	}

	err := h.Storage.CreateBucket(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(storage.BucketResponse{
		Name:    name,
		Success: true,
		Action:  "create",
	})
}

func (h *Handler) deleteBucket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty Bucket Name", http.StatusBadRequest)
		return
	}

	err := h.Storage.RemoveBucket(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(storage.BucketResponse{
		Name:    name,
		Success: true,
		Action:  "delete",
	})
}

func (h *Handler) listItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty Bucket Name", http.StatusBadRequest)
		return
	}

	items, err := h.Storage.ListItems(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bucket := r.URL.Query().Get("bucket")
	filename := r.URL.Query().Get("filename")
	if bucket == "" || filename == "" {
		http.Error(w, "You must supply a bucket (bucket) name and file name (filename)", http.StatusBadRequest)
		return
	}

	err := h.Storage.UploadFile(bucket, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(storage.ItemResponse{
		Bucket:  bucket,
		File:    filename,
		Success: true,
		Action:  "upload",
	})
}
