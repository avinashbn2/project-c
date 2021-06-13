package main

import (
	"cproject/internal/config"
	"cproject/internal/resource"
	"cproject/internal/user"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db   *sqlx.DB
	cfg  *config.Config
	Addr string
}

func NewServer(port string, db *sqlx.DB, cfg *config.Config) *Server {
	return &Server{
		Addr: port,
		db:   db,
		cfg:  cfg,
	}
}
func AllowOriginFunc(r *http.Request, origin string) bool {

	return true
}

func (s *Server) Handler() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Set-Cookie", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	userService := user.NewService(user.NewRepo(s.db))
	authHandlers := user.RegisterAuthHandlers(s.cfg, userService)
	resourceRepo := resource.NewRepo(s.db)
	resourceHandlers := resource.RegisterHandlers(s.cfg, resource.NewService(resourceRepo))
	userHandlers := user.RegisterHandlers(userService)
	router.Route("/health_check", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("test"))
		})
	})

	router.Route("/auth", authHandlers)
	router.Group(func(r chi.Router) {
		r.Use(user.IsLoggedIn(s.cfg))
		r.Route("/user", userHandlers)
		r.Route("/resource", resourceHandlers)
	})

	return router

}
