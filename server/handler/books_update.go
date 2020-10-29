package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jserrano27/bookStore/db"
	"github.com/jserrano27/bookStore/server/responser"
)

func UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.Atoi(chi.URLParam(r, "id"))
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

		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(&body)

		book := generateBookForUpdate(body)
		/*ID:        id,
			Title:     body["title"].(string),
			Author:    body["author"].(string),
			Year:      int(body["year"].(float64)),
			UpdatedAt: time.Now(),
		}*/

		res, err := db.UpdateBook(conn, book)

		if err != nil {
			errMsg := responser.NewErrorResponse(
				http.StatusInternalServerError,
				err.Error(),
			)
			render.Render(w, r, errMsg)
			return
		}

		fmt.Println(res)

		//TODO return info of updated book
	}

}

func generateBookForUpdate(body map[string]interface{}) *db.BookModel {
	book := &db.BookModel{}

	return book
}
