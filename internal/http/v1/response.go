package v1

import (
	"encoding/json"
	"net/http"
	"social/internal/domain"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type getPostsResponse struct {
	Data []domain.Post `json:"data"`
}

func newErrorResponse(w http.ResponseWriter, msg string, status int) {
	resp, _ := json.Marshal(errorResponse{msg}) //nolint:errcheck
	w.WriteHeader(status)
	w.Write(resp) //nolint:errcheck
}

func newStatusReponse(w http.ResponseWriter, msg string, status int) {
	resp, _ := json.Marshal(statusResponse{msg}) //nolint:errcheck
	w.WriteHeader(status)
	w.Write(resp) //nolint:errcheck
}

func newGetPostsResponse(w http.ResponseWriter, products []domain.Post, status int) {
	resp, _ := json.Marshal(getPostsResponse{products}) //nolint:errcheck
	w.WriteHeader(status)
	w.Write(resp) //nolint:errcheck
}
