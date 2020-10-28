package main

import (
	"fmt"
	"net/http"

	"github.com/jserrano27/bookStore/server/routes"
)

func main() {
	port := ":3000"

	server := routes.SetUp()

	fmt.Println("Serving on http://localhost" + port)
	http.ListenAndServe(port, server)
}
