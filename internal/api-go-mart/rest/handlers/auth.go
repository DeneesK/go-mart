package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/DeneesK/go-mart/internal/api-go-mart/entities"
	"github.com/DeneesK/go-mart/internal/api-go-mart/utils/jwt"
)

const (
	requestTimeout = 1 * time.Second
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	requestContext, requestCancel := context.WithTimeout(r.Context(), requestTimeout)
	defer requestCancel()

	var user entities.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.db.Register(requestContext, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.db.Login(requestContext, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateToken(user.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Header().Add("Authorization", "Bearer"+token)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	requestContext, requestCancel := context.WithTimeout(r.Context(), requestTimeout)
	defer requestCancel()

	var user entities.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.db.Login(requestContext, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateToken(user.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Header().Add("Authorization", "Bearer"+token)
	w.WriteHeader(http.StatusOK)
}
