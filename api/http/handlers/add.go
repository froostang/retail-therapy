package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/api/product"
)

// Define the AddProduct struct
type AdderData struct {
	URL string `json:"url"`
}

type addRequest struct {
	URL string
}

func (ad *AdderData) ToRequest() *addRequest {
	// validate url
	return &addRequest{
		URL: html.EscapeString(ad.URL),
	}
}

type Adder struct {
	cache  map[string]product.Scraped
	logger logging.Logger
}

func (a *Adder) insert(url string, p product.Scraped) {
	a.cache[url] = p
}

func NewAdderManager(logger logging.Logger) *Adder {
	return &Adder{
		logger: logger,
		cache:  make(map[string]product.Scraped, 100),
	}
}

func (sm *Adder) AdderHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON body into an AddProduct struct
	var addData AdderData
	if err := json.Unmarshal(body, &addData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	ps := product.Scraped

	// Respond to the client
	responseMessage := fmt.Sprintf("Received your message: %s", ps)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
}
