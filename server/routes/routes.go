package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jserrano27/bookStore/platform/books"
	"github.com/jserrano27/bookStore/server/handler"
	"github.com/jserrano27/bookStore/server/utils"
)

func SetUp() *chi.Mux {
	store := books.NewStore()
	counter := utils.NewCounter()
	r := chi.NewRouter()
	r.MethodNotAllowed(handler.MethodNotAllowed())

	r.Use(
		middleware.SetHeader("content-type", "application/json"),
		middleware.Logger,
		middleware.Recoverer,
		middleware.RealIP,
	)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks(store))
		r.Get("/{id}", handler.GetOneBook(store))
		r.Delete("/{id}", handler.DeleteBook(store))
	})

	r.Post("/books", handler.AddBook(store, counter))

	return r
}
