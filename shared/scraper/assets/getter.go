package assets

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Getter struct{}

func (g Getter) GetImage(body []byte) (string, error) {
	var imgURL string
	r := bytes.NewReader(body)

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return imgURL, fmt.Errorf("goquery failed: %w", err)
	}

	// Extract image URL
	imageURL, exists := doc.Find("img").Attr("src")
	if !exists {
		return imgURL, errors.New("image source doesn't exist")
	}

	imgURL = imageURL
	return imgURL, nil
}

func (g Getter) GetPrice(body []byte) (float64, error) {

	// FIXME: kind of hacky, scraping can be difficult... consider escaping
	re := regexp.MustCompile(`"current_retail\\":(\d+\.\d+)`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return 0, fmt.Errorf("price not found")
	}

	priceStr := matches[1]
	var price float64
	_, err := fmt.Sscanf(priceStr, "%f", &price)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (g Getter) GetName(body []byte) (string, error) {

	var name string
	r := bytes.NewReader(body)

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return name, fmt.Errorf("name goquery failed: %w", err)
	}

	// Extract image URL
	name = doc.Find("#pdp-product-title-id").Text()
	if name == "" {
		return name, errors.New("name source doesn't exist")
	}

	return name, nil
}

func (g Getter) GetDescription(body []byte) (string, error) {

	var desc string
	r := bytes.NewReader(body)

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return desc, fmt.Errorf("name goquery failed: %w", err)
	}

	// Extract image URL
	desc = doc.Find(`[data-test="item-details-description"]`).Text()
	if desc == "" {
		return desc, errors.New("name source doesn't exist")
	}

	return desc, nil
}
