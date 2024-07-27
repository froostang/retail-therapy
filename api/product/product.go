package product

import "github.com/froostang/retail-therapy/shared/scraper"

type Requester interface {
	Scrape(string) Scraped
}

type ScraperRequest struct {
}

func (s *ScraperRequest) Scrape(url string) (*Scraped, error) {
	i, p, err := scraper.ScrapeForImagePrice(url)
	if err != nil {
		return nil, err
	}
	return &Scraped{
		ImageURL: i,
		Price:    p,
		Name:     "test",
	}, nil
}

type Scraped struct {
	ImageURL    string
	Name        string
	Description string
	Price       string
}
