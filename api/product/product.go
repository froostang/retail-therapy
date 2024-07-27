package product

type Requester interface {
	Scrape(string) Scraped
}

type Scraper struct {
}

func (s *Scraper) Scrape(url string) Scraped {
	i, p, err := s.ScrapeForImagePrice(url)
}

type Scraped struct {
	ImageURL    string
	Name        string
	Description string
	Price       string
}
