package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jserrano27/bookStore/server/routes"
)

func main() {
	server := routes.SetUp()
	port := ":3005"

	fmt.Println("Serving on http://localhost" + port)
	if err := http.ListenAndServe(port, server); err != nil {
		log.Printf("Failed to start server: %v.\n", err)
		os.Exit(100)
	}
}
