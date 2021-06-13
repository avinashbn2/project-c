// Package main provides ...
package resource

import (
	"cproject/internal/config"
	"cproject/internal/models"
	"cproject/internal/user"
	"cproject/pkg/errors"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`          // user-level status message
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		StatusText: "Error rendering response.",
		ErrorText:  err.Error(),
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type handler struct {
	srv *service
}
type QueryParams struct {
	page      int
	limit     int
	sortOrder string
	sortParam string
	offset    int
	search    string
}

type QueryWithUser struct {
	user *models.User
	QueryParams
}

func RegisterHandlers(cfg *config.Config, srv *service) func(chi.Router) {
	h := &handler{srv}
	return func(r chi.Router) {

		r.Get("/", h.get())
		r.Get("/{id}", h.getByID())
		r.Get("/user", h.getByUserLikes())
		r.Get("/trending", h.getByTrending())
		r.Post("/", h.create())
		r.Put("/{id}", h.update())
	}
}

func (h *handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resource := &ResourceItem{}
		err := json.NewDecoder(r.Body).Decode(resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := h.srv.create(resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		rsp, err := json.Marshal(response)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		w.Write([]byte(rsp))
	}

}

func (h *handler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			//w.Header().Set("Content-type", "application/json")
			//w.WriteHeader(405)
			//render.Render(w, r, ErrRender(err))
			limit = 20
			return
		}
		sortBy := r.URL.Query().Get("sortBy")
		var sortOrder = "desc"
		var sortParam = "created_at"
		if sortBy != "" {

			sortPair := strings.Split(sortBy, ":")

			sortOrder = sortPair[1]
			sortParam = sortPair[0]
		}
		search := r.URL.Query().Get("search")
		offset := (page - 1) * limit
		user, err := user.GetUser(r.Context())
		if err != nil {
			log.Println("HANDLER", err)
			render.Render(w, r, errors.ErrNotAuthorized)

			return
		}

		response, err := h.srv.get(&QueryWithUser{user, QueryParams{page, limit, sortOrder, sortParam, offset, search}})
		if err != nil {
			log.Println("HANDLER", err)
		}
		bytes, err := json.Marshal(&response)
		if err != nil {
			log.Println("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))

	}
}

func (h *handler) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Fatal(err)
		}

		data, err := h.srv.getByID(id)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := json.Marshal(&data)
		if err != nil {
			log.Fatal("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))
	}
}

func (h *handler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resource := &ResourceItem{}
		err := json.NewDecoder(r.Body).Decode(resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := chi.URLParam(r, "id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		resource.ID = idInt

		data, err := h.srv.update(resource)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := json.Marshal(&data)
		if err != nil {
			log.Fatal("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))
	}

}

func (h *handler) getByUserLikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			//w.Header().Set("Content-type", "application/json")
			//w.WriteHeader(405)
			//render.Render(w, r, ErrRender(err))
			limit = 20
			return
		}
		sortBy := r.URL.Query().Get("sortBy")
		var sortOrder = "desc"
		var sortParam = "created_at"
		if sortBy != "" {

			sortPair := strings.Split(sortBy, ":")

			sortOrder = sortPair[1]
			sortParam = sortPair[0]
		}
		search := r.URL.Query().Get("search")
		offset := (page - 1) * limit
		user, err := user.GetUser(r.Context())
		if err != nil {
			log.Println("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		response, err := h.srv.getByUserLikes(&QueryWithUser{user, QueryParams{page, limit, sortOrder, sortParam, offset, search}})
		if err != nil {
			log.Println("HANDLER", err)
		}
		bytes, err := json.Marshal(&response)
		if err != nil {
			log.Println("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))

	}
}

func (h *handler) getByTrending() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 20
			return
		}
		sortBy := r.URL.Query().Get("sortBy")
		var sortOrder = "desc"
		var sortParam = "created_at"
		if sortBy != "" {

			sortPair := strings.Split(sortBy, ":")
			sortOrder = sortPair[1]
			sortParam = sortPair[0]
		}
		search := r.URL.Query().Get("search")
		offset := (page - 1) * limit
		user, err := user.GetUser(r.Context())
		if err != nil {
			log.Println("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		response, err := h.srv.getByTrending(&QueryWithUser{user, QueryParams{page, limit, sortOrder, sortParam, offset, search}})
		if err != nil {
			log.Println("HANDLER", err)
		}
		bytes, err := json.Marshal(&response)
		if err != nil {
			log.Println("HANDLER", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bytes))

	}
}
