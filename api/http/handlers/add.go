package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/froostang/retail-therapy/api/cache"
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
	cache     *cache.Products
	logger    logging.Logger
	requester product.Requester
}

func (a *Adder) insert(url string, p product.Scraped) {
	a.cache.Insert(url, p)
}

func NewAdderManager(logger *loggers.ZapLogger, c *cache.Products) *Adder {
	return &Adder{
		logger:    logger,
		cache:     c,
		requester: product.NewScrapeRequester(logger),
	}
}

func (sm *Adder) AdderHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var requestData AdderData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, fmt.Errorf("failed to parse request body: %w", err).Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ar := requestData.ToRequest()

	// TODO: scrape function on ar.URL
	// insert
	sp, err := sm.requester.Scrape(ar.URL)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to scrape product: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	sm.insert(ar.URL, *sp)

	// Respond to the client
	responseMessage := fmt.Sprintf("Received your message: %s", ar.URL)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
}
