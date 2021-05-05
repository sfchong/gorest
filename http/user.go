package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sfchong/gorest"
	"net/http"
	"strconv"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")
	fullname := r.URL.Query().Get("fullname")

	if email == "" && fullname == "" {
		json.NewEncoder(w).Encode(nil)
		return
	}

	filter := gorest.UserFilter{
		FullName: fullname,
		Email:    email,
	}

	users, err := s.UserService.Get(filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (s *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	user, err := s.UserService.GetById(id)

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user gorest.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.UserService.Create(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	var user gorest.UserUpdate

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.UserService.Update(id, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	err = s.UserService.DeleteById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
