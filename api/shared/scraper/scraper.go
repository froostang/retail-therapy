package scraper

import (
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
)

var ErrBadURL = errors.New("bad URL")

func isTargetURL(str string) bool {
	u, err := url.Parse(str)
	fmt.Println(u.Hostname())
	return err == nil && u.Scheme != "" && u.Host != "" && u.Host == "www.target.com"
}

type AssetGetter interface {
	GetImage(body []byte) (string, error)
	GetPrice(body []byte) (float64, error)
	GetName(body []byte) (string, error)
	GetDescription(body []byte) (string, error)
}

type Result struct {
	Image       string
	Price       string
	Name        string
	Description string
}

// ScrapeForImagePrice uses AssetGetter to fetch image and price
func Scrape(url string, getter AssetGetter) (Result, error) {

	result := Result{}

	if !isTargetURL(url) {
		return result, ErrBadURL
	}

	sanitized := html.EscapeString(url)

	resp, err := http.Get(sanitized)
	if err != nil {
		return result, ErrBadURL
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, errors.New("can't read body")
	}

	i, err := getter.GetImage(bodyBytes)
	if err != nil {
		return result, fmt.Errorf("could not get image: %w", err)
	}
	result.Image = i

	p, err := getter.GetPrice(bodyBytes)
	if err != nil {
		return result, fmt.Errorf("could not get price: %w", err)
	}
	result.Price = fmt.Sprintf("%.2f", p)

	n, err := getter.GetName(bodyBytes)
	if err != nil {
		return result, fmt.Errorf("could not get name: %w", err)
	}
	result.Name = n

	d, err := getter.GetDescription(bodyBytes)
	if err != nil {
		return result, fmt.Errorf("could not get description: %w", err)
	}
	result.Description = d

	return result, nil
}
