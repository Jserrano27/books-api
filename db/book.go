package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

type BookModel struct {
	tableName struct{}  `pg:"books"`
	ID        int       `pg:"id,unique,pk"`
	Title     string    `pg:"title,notnull"`
	Author    string    `pg:"author,notnull"`
	Year      int       `pg:"year,notnull"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func CreateBookTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.Model(&BookModel{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Failed to create book table: %v.\n", createErr)
		return createErr
	}

	return nil
}

func GetOneBook(db *pg.DB, id int) *Book {
	book := &Book{ID: id}

	err := db.Model(book).WherePK().Select()

	if err != nil {
		fmt.Printf("Failed getting the book: %v.\n", err)
		return nil
	}

	return book
}

func GetAllBooks(db *pg.DB) *[]Book {
	books := &[]Book{}

	err := db.Model(books).Order("id ASC").Select()

	if err != nil {
		fmt.Printf("Failed getting the book: %v.\n", err)
		return nil
	}

	return books
}

func InsertOneBook(db *pg.DB, book *BookModel) (*BookModel, error) {
	_, err := db.Model(book).Insert()

	if err != nil {
		return nil, errors.New("error inserting book")
	}

	return book, nil
}

func DeleteBook(db *pg.DB, id int) error {
	book := &Book{ID: id}
	_, err := db.Model(book).WherePK().Delete()

	if err != nil {
		return errors.New("error deleting book")
	}

	return nil
}

func UpdateBook(db *pg.DB, book *BookModel) (orm.Result, error) {
	res, err := db.Model(book).WherePK().Update()

	if err != nil {
		return nil, errors.New("error updating book: " + err.Error())
	}

	return res, nil
}
