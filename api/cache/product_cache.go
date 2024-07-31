package cache

import (
	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/api/product"
)

// client API
type ProductCacher interface {
	Insert(string, product.Scraped)
	Get(string) product.Scraped
	GetAll() []product.Scraped
}

// TODO: This cache stuff should live in shared module
// this should/could then be a simple API that shared.Cache implements
// or an internal representation that wraps the more generic cache

// TODO: maybe move to product pkg?
type Products struct {
	logger logging.Logger
	values []cacheVal
	max    int
}

type cacheVal struct {
	url     string
	product product.Scraped
}

func NewForProducts(logger logging.Logger, max int) *Products {
	m := 100

	// hard cap max
	if max < 10000 {
		m = max
	}

	return &Products{
		logger: logger,
		max:    m,
		values: make([]cacheVal, 0, m),
	}
}

func (p *Products) Insert(url string, ps product.Scraped) {
	if len(p.values) >= p.max {
		p.values = p.values[p.max/2:] // cut the first half
	}

	p.values = append(p.values, cacheVal{url: url, product: ps})

	p.logger.Info("cache values:")

	for _, item := range p.values {
		p.logger.Info(item.url)
	}
}

func (p *Products) GetAll() []product.Scraped {
	scraped := make([]product.Scraped, 0, len(p.values))
	for _, c := range p.values {
		scraped = append(scraped, c.product)
	}

	return scraped
}

func (p *Products) Get(key string) product.Scraped {
	for _, c := range p.values {
		if c.product.Name == key {
			return c.product
		}
	}

	return product.Scraped{}
}
