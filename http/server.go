package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/sfchong/gorest"
	"net/http"
)

type Server struct {
	router *chi.Mux

	UserService gorest.UserService
}

func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
	}
	s.routes(s.router)

	return s
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", s.router)
}
