package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/api/product"
	"github.com/froostang/retail-therapy/shared/loggers"
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

type cacheVal struct {
	url     string
	product product.Scraped
}

type Adder struct {
	cache  []cacheVal
	logger logging.Logger
}

func (a *Adder) insert(url string, p product.Scraped) {
	if len(a.cache) >= 100 {
		a.cache = a.cache[50:]
	}

	a.cache = append(a.cache, cacheVal{url: url, product: p})
}

func NewAdderManager(logger *loggers.ZapLogger) *Adder {
	return &Adder{
		logger: logger,
		cache:  make([]cacheVal, 0, 100),
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

	ar := addData.ToRequest()

	// TODO: scrape function on ar.URL
	// insert
	sm.insert(ar.URL, product.Scraped{Name: "test"})

	// Respond to the client
	responseMessage := fmt.Sprintf("Received your message: %s", ar.URL)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
}
