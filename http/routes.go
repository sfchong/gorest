package http

import "github.com/go-chi/chi/v5"

func (s *Server) routes(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/", s.handleGetUser)
		r.Post("/", s.handleCreateUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", s.handleGetUserById)
			r.Put("/", s.handleUpdateUserById)
			r.Delete("/", s.handleDeleteUserById)
		})
	})
}
