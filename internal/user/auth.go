package user

import (
	"cproject/internal/config"
	"cproject/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// TODO randomize it
var randomState = "random"

type Claims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}

type gAuthResponse struct {
	Email string `json:"email"`
}

type auth struct {
	cfg         *config.Config
	gauthConfig *oauth2.Config
	userService Service
}

func RegisterAuthHandlers(cfg *config.Config, service Service) func(chi.Router) {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  cfg.SERVER_URL + "/auth/callback",
		ClientID:     cfg.GAUTH_CLIENT_ID,
		ClientSecret: cfg.GAUTH_CLIENT_SECRET,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	authHandler := &auth{cfg, googleOauthConfig, service}
	return func(router chi.Router) {
		router.Get("/login", authHandler.login())
		router.Get("/logout", authHandler.logout())
		router.Get("/callback", authHandler.callback())
	}
}

func (a *auth) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := r.URL.Query().Get("provider")
		if provider == "google" {
			url := a.gauthConfig.AuthCodeURL(randomState)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			fmt.Fprintf(w, "Provider not supported")
			return
		}

	}
}

// oauth callback handler
func (a *auth) callback() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("state") != randomState {
			fmt.Fprintf(w, "State is not valid")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return

		}

		gtoken, err := a.gauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
		if err != nil {
			fmt.Fprintf(w, "Couldnt fetch token %s", err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", gtoken.AccessToken))
		if err != nil {

			fmt.Fprintf(w, "couldn't perform get request")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)

		if err != nil {

			fmt.Fprintf(w, "couldn't parse response")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
		resp := &gAuthResponse{}

		err = json.Unmarshal(content, resp)
		if err != nil {

			fmt.Fprintf(w, "failed to unmarshal %v", err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

		var userID int
		user := &models.User{Email: resp.Email}
		registeredUser, err := a.userService.IsRegisteredUser(user)
		if err != nil {
			fmt.Println("error occurred", err)
		}
		if registeredUser == nil {

			newUser, err := a.userService.RegisterUser(user)
			if err != nil {
				fmt.Println("failed to registed user", err)
			} else {
				userID = newUser.ID
				fmt.Println("new registered", newUser)
			}
		} else {

			userID = registeredUser.ID
			fmt.Println("user is registed", registeredUser)
		}
		// 3 days
		expirationTime := time.Now().Add(24 * 3 * time.Hour)
		claims := &Claims{
			Email: resp.Email,
			ID:    userID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(a.cfg.JWTKEY))
		if err != nil {
			fmt.Fprintf(w, "Failed to sign string")
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "Token",
			Value:    tokenString,
			Expires:  expirationTime,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			//HttpOnly: true,
			//Secure: true,
		})

		http.Redirect(w, r, a.cfg.CLIENT_URL+"/resources", http.StatusTemporaryRedirect)
	}
}

func (a *auth) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		http.SetCookie(w, &http.Cookie{
			Name:   "Token",
			Value:  "rubbish",
			MaxAge: -1,
			Path:   "/",
			//HttpOnly: true,
			//Secure: true,
		})
		http.Redirect(w, r, a.cfg.CLIENT_URL, http.StatusTemporaryRedirect)
	}
}
