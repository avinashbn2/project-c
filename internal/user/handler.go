package user

import (
	"cproject/internal/config"
	"cproject/internal/models"
	"cproject/pkg/errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type handler struct {
	srv *service
}

func RegisterHandlers(service *service) func(chi.Router) {
	h := &handler{srv: service}
	return func(router chi.Router) {
		router.Post("/like", h.updateLikes())
		router.Get("/like", h.getLikes())
		router.Get("/readingList", h.getReadingList())
		router.Post("/readingList", h.addToReadingList())
		router.Post("/view", h.updateViews())
		router.Get("/view", h.getViews())

	}

}
func (uh *handler) updateLikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userResource := &UserResource{}
		err := json.NewDecoder(r.Body).Decode(userResource)
		defer r.Body.Close()
		fmt.Println("Context", r.Context())
		user, err := GetUser(r.Context())
		if err != nil {

			render.Render(w, r, errors.ErrNotAuthorized)
			return
		}

		id := user.ID

		userResource.Uid = id

		err = uh.srv.updateLikes(userResource)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"msg":"success"}`))

	}
}

func (uh *handler) getLikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get likes"))

	}
}

func (uh *handler) getReadingList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get reasing list"))
	}
}

func (uh *handler) addToReadingList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh *handler) updateViews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh *handler) getViews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func IsLoggedIn(cfg *config.Config) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("INISIDE MIDDLEWARE")

			for _, cookie := range r.Cookies() {
				fmt.Println("cppokero", cookie.Name)
			}

			c, err := r.Cookie("Token")
			if err != nil {

				if err == http.ErrNoCookie {

					w.WriteHeader(http.StatusUnauthorized)
					next.ServeHTTP(w, r)
				}

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			tokenString := c.Value

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTKEY), nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					log.Fatal(err.Error())
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				log.Fatal(err.Error())
				return
			}

			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				log.Fatal(err.Error())
				return
			}

			user := &models.User{ID: claims.ID}
			next.ServeHTTP(w, r.WithContext(withUser(r.Context(), user)))
		})
	}
}
