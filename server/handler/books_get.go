package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jserrano27/bookStore/server/responser"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
	"github.com/jserrano27/bookStore/platform/books"
)

func GetAllBooks(store books.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := store.GetAllBooks()
		json.NewEncoder(w).Encode(books)
	}
}

func GetOneBook(store books.Getter) http.HandlerFunc {
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

		book, err := store.GetOneBook(id)
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
