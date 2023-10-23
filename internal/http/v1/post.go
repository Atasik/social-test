package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"social/internal/domain"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", appJSON)
	if r.Header.Get("Content-Type") != appJSON {
		newErrorResponse(w, "unknown payload", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, "server error", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	var post domain.Post
	if err = json.Unmarshal(body, &post); err != nil {
		newErrorResponse(w, "cant unpack payload", http.StatusBadRequest)
		return
	}

	if err = h.validator.Struct(post); err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.ID, err = h.services.Create(post)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Post was created: %d", post.ID)

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(post); err != nil {
		newErrorResponse(w, "server error", http.StatusInternalServerError)
	}
}

func (h *handler) getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", appJSON)

	posts, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newGetPostsResponse(w, posts, http.StatusOK)
}

func (h *handler) getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", appJSON)

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["postId"])
	if err != nil {
		newErrorResponse(w, "Bad Id", http.StatusBadRequest)
		return
	}

	post, err := h.services.GetByID(postID)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(post); err != nil {
		newErrorResponse(w, "server error", http.StatusInternalServerError)
	}

}

func (h *handler) deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", appJSON)

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["postId"])
	if err != nil {
		newErrorResponse(w, "Bad Id", http.StatusBadRequest)
		return
	}

	if err = h.services.Delete(postID); err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Infof("Post was created: %d", postID)

	newStatusReponse(w, "done", http.StatusOK)
}

func (h *handler) updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", appJSON)
	if r.Header.Get("Content-Type") != appJSON {
		newErrorResponse(w, "unknown payload", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["postId"])
	if err != nil {
		newErrorResponse(w, "Bad Id", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, "server error", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	var input domain.UpdatePostInput
	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, "cant unpack payload", http.StatusBadRequest)
		return
	}

	if err = input.Validate(); err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.services.Update(postID, input); err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Infof("Post was updated: %d", postID)

	post, err := h.services.GetByID(postID)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(post); err != nil {
		newErrorResponse(w, "server error", http.StatusInternalServerError)
	}
}
