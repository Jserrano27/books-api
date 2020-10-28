package books

import (
	"errors"
)

type Getter interface {
	GetOneBook(id int) (Book, error)
	GetAllBooks() []Book
}

type Adder interface {
	AddBook(book Book)
}

type Deleter interface {
	DeleteBook(id int) (Book, error)
}

type Store struct {
	Books []Book
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func NewStore() *Store {
	return &Store{
		Books: []Book{},
	}
}

func (s *Store) AddBook(book Book) {
	s.Books = append(s.Books, book)
}

func (s *Store) DeleteBook(id int) (Book, error) {
	book := Book{}

	for i, v := range s.Books {
		if v.ID == id {
			book = s.Books[i]
			s.Books = s.Books[:i+copy(s.Books[i:], s.Books[i+1:])]
			return book, nil
		}
	}

	return book, errors.New("book not found with the provided id")
}

func (s *Store) GetOneBook(id int) (Book, error) {
	book := Book{}

	for _, v := range s.Books {
		if v.ID == id {
			book = v
			return book, nil
		}
	}

	return book, errors.New("book not found with the provided id")
}

func (s *Store) GetAllBooks() []Book {
	return s.Books
}
