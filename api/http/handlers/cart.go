package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// CartData represents the data structure for the cart item
type CartData struct {
	Name string `json:"name"`
}

// CartHandler processes the addition of an item to the cart
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

	// Insert item into cart (assuming this function updates the cart in some way)
	sm.cart.Insert(decodedName, sm.cache.Get(decodedName))

	// Get the updated cart count
	cartCount := len(sm.cart.GetAll())

	// Respond with the updated cart count as JSON
	response := map[string]int{"count": cartCount}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response as JSON: %v", err), http.StatusInternalServerError)
	}
}
