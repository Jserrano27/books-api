package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jserrano27/bookStore/server/responser"
)

func MethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errMsg := responser.NewErrorResponse(
			http.StatusMethodNotAllowed,
			"method not allowed",
		)
		render.Render(w, r, errMsg)
	}
}
