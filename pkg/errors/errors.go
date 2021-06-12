package errors

import (
	"net/http"

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

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Not found."}
var ErrNotAuthorized = &ErrResponse{HTTPStatusCode: 401, StatusText: "Not authorized"}
