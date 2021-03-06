package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/jserrano27/bookStore/db"
	"github.com/jserrano27/bookStore/server/responser"
	"github.com/jserrano27/bookStore/server/utils"
)

func AddOneBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(&body)

		err := utils.ValidatePayload(body)
		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusBadRequest,
				err.Error(),
			)
			render.Render(w, r, errMsg)
			return
		}

		book := &db.BookModel{
			Title:     body["title"].(string),
			Author:    body["author"].(string),
			Year:      int(body["year"].(float64)),
			CreatedAt: time.Now(),
		}

		conn := db.Connect()
		defer db.CloseConnection(conn)

		res, err := db.InsertOneBook(conn, book)
		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusInternalServerError,
				err.Error(),
			)
			render.Render(w, r, errMsg)
			return
		}

		json.NewEncoder(w).Encode(db.Book{
			ID:     res.ID,
			Title:  res.Title,
			Author: res.Author,
			Year:   res.Year,
		})
	}
}
