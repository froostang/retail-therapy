package handlers

import (
	"fmt"
	"html"
	"net/http"
)

type CartData struct {
	Name string `json:"name"`
}

func (sm *ShoppingManager) CartHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse form data: %v", err), http.StatusBadRequest)
		return
	}

	// Get the name value from the form data
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Missing 'name' field in form data", http.StatusBadRequest)
		return
	}

	// Decode HTML entities
	decodedName := html.UnescapeString(name)

	sm.logger.Info(decodedName)

	sm.cart.Insert(decodedName, sm.cache.Get(decodedName))

}
