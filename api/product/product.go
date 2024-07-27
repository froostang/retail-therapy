package product

import (
	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/shared/scraper"
	"github.com/froostang/retail-therapy/shared/scraper/assets"
)

type Requester interface {
	Scrape(string) (*Scraped, error)
}

type ScrapeRequester struct {
	logger logging.Logger
}

func NewScrapeRequester(logger logging.Logger) *ScrapeRequester {
	return &ScrapeRequester{
		logger: logger,
	}
}

func (s *ScrapeRequester) Scrape(url string) (*Scraped, error) {

	s.logger.Info(url)
	i, p, err := scraper.ScrapeForImagePrice(url, assets.Getter{})
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
