package main

import (
	"fmt"
	"net/http"

	"github.com/cmurray4492/go-project1/pkg/handlers"
)

const portNumber = ":8080"

// main application entry point
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
