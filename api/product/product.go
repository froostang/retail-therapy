package product

import (
	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/api/shared/scraper"
	"github.com/froostang/retail-therapy/api/shared/scraper/assets"
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
	result, err := scraper.Scrape(url, assets.Getter{})
	if err != nil {
		return nil, err
	}
	return &Scraped{
		ImageURL:    result.Image,
		Price:       result.Price,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

type Scraped struct {
	ImageURL    string
	Name        string
	Description string
	Price       string
}
