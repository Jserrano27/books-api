package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jserrano27/bookStore/server/handler"
)

func SetUp() *chi.Mux {
	r := chi.NewRouter()
	r.MethodNotAllowed(handler.MethodNotAllowed())

	r.Use(
		middleware.SetHeader("content-type", "application/json"),
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks())
		r.Get("/{id}", handler.GetOneBook())
		r.Put("/{id}", handler.UpdateBook())
		r.Delete("/{id}", handler.DeleteBook())
	})

	r.Post("/books", handler.AddOneBook())

	return r
}
