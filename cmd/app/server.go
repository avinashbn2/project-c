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
	router.Route("/resource", func(r chi.Router) {
		r.Get("/", s.resourceRepo.FindAll())
		r.Get("/{id}", s.resourceRepo.FindByID())
		r.Post("/", s.resourceRepo.Add())
		r.Put("/{id}", s.resourceRepo.Update())
		// r.Delete(/, s.resourceRepo.Remove())

	})
	return router

}
