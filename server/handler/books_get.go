package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jserrano27/bookStore/db"

	"github.com/jserrano27/bookStore/server/responser"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

func GetAllBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := db.Connect()
		defer db.CloseConnection(conn)

		books := db.GetAllBooks(conn)
		json.NewEncoder(w).Encode(books)
	}
}

func GetOneBook() http.HandlerFunc {
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

		if book == nil {
			errMsg := responser.NewErrorResponse(
				http.StatusNotFound,
				"book not found with the provided id",
			)
			render.Render(w, r, errMsg)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}
