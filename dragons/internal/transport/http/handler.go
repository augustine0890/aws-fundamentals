package http

import (
	"dragons/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	h := &Handler{
		Storage: storage,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	s3 := h.Router.PathPrefix("/s3").Subrouter()
	s3.HandleFunc("/buckets", h.listBuckets)
	s3.HandleFunc("/bucket/items", h.listItems)
	s3.HandleFunc("/bucket", h.createBucket).Methods(http.MethodPost)
	s3.HandleFunc("/bucket", h.queryBucket).Methods(http.MethodGet)
	s3.HandleFunc("/bucket", h.deleteBucket).Methods(http.MethodDelete)

	// auth := h.Router.PathPrefix("/auth").Subrouter()
	// auth.HandleFunc("/", app.testing)

}
