package main

import (
	"fmt"
	"net/http"

	"github.com/froostang/retail-therapy/shared/middleware"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Apply JWT middleware to all routes
	http.Handle("/", middleware.JWTMiddleware(http.DefaultServeMux))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
