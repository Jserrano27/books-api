package db

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	var db = pg.Connect(&pg.Options{
		Addr:         "localhost:5432",
		Database:     "books_db",
		User:         "postgres",
		Password:     "0000",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		PoolSize:     20,
	})

	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}

	CreateBookTable(db)

	log.Printf("Connection to database successful.\n")

	return db
}

func CloseConnection(db *pg.DB) {
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing the connection: %v.\n", closeErr)
	}

	log.Println("Connection closed succesfully")
}
