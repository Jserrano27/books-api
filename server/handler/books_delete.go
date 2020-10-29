package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jserrano27/bookStore/db"
	"github.com/jserrano27/bookStore/server/responser"
)

func DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusBadRequest,
				"ID must be a number",
			)
			render.Render(w, r, errMsg)
			return
		}

		conn := db.Connect()
		defer db.CloseConnection(conn)

		book := db.GetOneBook(conn, id)

		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusNotFound,
				err.Error(),
			)
			render.Render(w, r, errMsg)
			return
		}

		err = db.DeleteBook(conn, id)
		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusNotFound,
				err.Error(),
			)
			render.Render(w, r, errMsg)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}
