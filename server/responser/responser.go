package responser

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	Success        bool   `json:"success"`
	Reason         string `json:"reason"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func NewErrorResponse(code int, message string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: code,
		Success:        false,
		Reason:         message,
	}
}
