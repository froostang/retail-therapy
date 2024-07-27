package assets

import (
	"bytes"
	"errors"

	"github.com/PuerkitoBio/goquery"
)

func GetImage(body []byte) (string, error) {
	var imgURL string
	r := bytes.NewReader(body)

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return imgURL, errors.New("can't make document")
	}

	// Extract image URL
	imageURL, exists := doc.Find("img").Attr("src")
	if exists {
		imgURL = imageURL
		return imgURL, errors.New("can't make document")
	} else {
		return imgURL, errors.New("image source doesn't exist")
	}
}
