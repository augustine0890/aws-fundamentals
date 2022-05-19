package http

import (
	"net/http"

	"dragons/internal/auth"
	"dragons/internal/storage"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Storage *storage.Storage
	Auth    *auth.Auth
}

func NewHandler(storage *storage.Storage, auth *auth.Auth) *Handler {
	h := &Handler{
		Storage: storage,
		Auth:    auth,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	s3 := h.Router.PathPrefix("/s3").Subrouter()
	s3.HandleFunc("/buckets", h.listBuckets)
	s3.HandleFunc("/bucket/items", h.listItems).Methods(http.MethodGet)
	s3.HandleFunc("/bucket/item", h.uploadFile).Methods(http.MethodPost)
	s3.HandleFunc("/bucket", h.createBucket).Methods(http.MethodPost)
	s3.HandleFunc("/bucket", h.queryBucket).Methods(http.MethodGet)
	s3.HandleFunc("/bucket", h.deleteBucket).Methods(http.MethodDelete)

	auth := h.Router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sigup", h.sigup).Methods(http.MethodPost)
	auth.HandleFunc("/confirm", h.confirm).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.login).Methods(http.MethodPost)
}
