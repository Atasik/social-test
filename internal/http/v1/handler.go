package v1

import (
	"net/http"
	"social/internal/service"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	appJSON = "application/json"
)

type handler struct {
	services  *service.Service
	logger    *zap.SugaredLogger
	validator *validator.Validate
}

func NewHandler(services *service.Service, logger *zap.SugaredLogger, validator *validator.Validate) *handler {
	return &handler{
		services:  services,
		logger:    logger,
		validator: validator,
	}
}

func (h *handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/post", h.createPost).Methods("POST")
	r.HandleFunc("/api/posts", h.getPosts).Methods("GET")
	r.HandleFunc("/api/post/{postId}", h.getPost).Methods("GET")
	r.HandleFunc("/api/post/{postId}", h.deletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{postId}", h.updatePost).Methods("PUT")

	mux := panicMiddleware(r)

	return mux
}
