package main

import (
	"cproject/cmd/repository"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	tagRepo      *repository.TagRepo
	resourceRepo *repository.ResourceRepo
	Addr         string
}

func NewServer(port string, repo *repository.ResourceRepo, tagRepo *repository.TagRepo) *Server {
	return &Server{
		Addr:         port,
		resourceRepo: repo,
		tagRepo:      tagRepo,
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

	router.Route("/tag", func(r chi.Router) {
		r.Post("/", s.tagRepo.Add())
		// r.Delete(/, s.resourceRepo.Remove())
	})

	return router

}
