package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jserrano27/bookStore/server/utils"

	"github.com/jserrano27/bookStore/platform/books"
)

func AddBook(store books.Adder, counter *utils.Counter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(&body)

		book := books.Book{
			ID:     counter.GetID(),
			Title:  body["title"].(string),
			Author: body["author"].(string),
			Year:   int(body["year"].(float64)),
		}

		store.AddBook(book)

		json.NewEncoder(w).Encode(book)
	}
}
