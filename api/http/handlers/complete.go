package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// CompleteHandler processes the completion of a purchase
func (sm *ShoppingManager) CompleteHandler(w http.ResponseWriter, r *http.Request) {
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

	// Extract form values
	email := r.FormValue("email")
	cardName := r.FormValue("card-name")
	cardNumber := r.FormValue("card-number")
	expiryDate := r.FormValue("expiry-date")
	cvv := r.FormValue("cvv")

	// Validate the extracted data (basic validation)
	if email == "" || cardName == "" || cardNumber == "" || expiryDate == "" || cvv == "" {
		http.Error(w, "All form fields are required", http.StatusBadRequest)
		return
	}

	// Decode HTML entities (if necessary)
	email = html.UnescapeString(email)
	cardName = html.UnescapeString(cardName)
	cardNumber = html.UnescapeString(cardNumber)
	expiryDate = html.UnescapeString(expiryDate)
	cvv = html.UnescapeString(cvv)
	sm.logger.Info(email, cardName, cardNumber, expiryDate, cvv)

	// Handle the payment processing and order completion logic
	// For simplicity, we'll assume success here
	// You should integrate with a payment gateway and order management system

	// Clear the cart after processing the order
	sm.cart.Clear() // Assuming there's a method to clear the cart

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Purchase completed successfully"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response as JSON: %v", err), http.StatusInternalServerError)
	}
}
