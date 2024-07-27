package scraper

import (
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"

	"github.com/froostang/retail-therapy/shared/scraper/assets"
)

var ErrBadURL = errors.New("bad URL")

func isTargetURL(str string) bool {
	u, err := url.Parse(str)
	fmt.Println(u.Hostname())
	return err == nil && u.Scheme != "" && u.Host != "" && u.Host == "www.target.com"
}

// only supports target URLs
func ScrapeForImagePrice(url string) (string, string, error) {

	var imageURL, price string

	if !isTargetURL(url) {
		return imageURL, price, fmt.Errorf("received: %s, %w", url, ErrBadURL)
	}

	sanitized := html.EscapeString(url)

	resp, err := http.Get(sanitized)
	if err != nil {
		return imageURL, price, fmt.Errorf("%w : %s", ErrBadURL, url)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return imageURL, price, errors.New("can't read body")
	}

	assets.GetImage(bodyBytes)

	i, err := assets.GetImage(bodyBytes)
	if err != nil {
		return imageURL, price, fmt.Errorf("could not get image: %w", err)
	}

	imageURL = i

	p, err := assets.GetPrice(bodyBytes)
	if err != nil {
		return imageURL, price, fmt.Errorf("could not get price: %w", err)
	}

	price = fmt.Sprintf("%.2f", p)
	return imageURL, price, nil
}
