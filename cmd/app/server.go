package main

import (
	"cproject/cmd/repository"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	resourceRepo *repository.ResourceRepo
	Addr         string
}

func NewServer(port string, repo *repository.ResourceRepo) *Server {
	return &Server{
		Addr:         port,
		resourceRepo: repo,
	}
}
func (s *Server) Handler() http.Handler {
	router := chi.NewRouter()
	router.Route("/test", func(r chi.Router) {
		r.Get("/", s.resourceRepo.FindAll())
	})
	return router
}
