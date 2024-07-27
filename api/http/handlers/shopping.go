package handlers

import (
	"fmt"
	"net/http"

	"github.com/froostang/retail-therapy/api/cache"
	"github.com/froostang/retail-therapy/api/logging"
	"github.com/froostang/retail-therapy/api/product"
	"github.com/froostang/retail-therapy/api/user"
)

type ShoppingData struct {
	User     user.User
	Products []product.Scraped
}

type ShoppingManager struct {
	logger logging.Logger
	cache  cache.ProductCacher
}

type optFunc func(sm *ShoppingManager)

func AddLogger(logger logging.Logger) optFunc {
	return func(sm *ShoppingManager) {
		sm.logger = logger
	}
}

func AddCacher(c cache.ProductCacher) optFunc {
	return func(sm *ShoppingManager) {
		sm.cache = c
	}
}

func NewShoppingManager(sm *ShoppingManager, opts ...optFunc) *ShoppingManager {
	for _, f := range opts {
		f(sm)
	}

	return sm
}

func (sm *ShoppingManager) ShoppingHandler(w http.ResponseWriter, r *http.Request) {
	user := user.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	if sm.logger == nil {
		panic("no logger")
	}

	products := sm.cache.Get()

	// TODO: fix directory structure issues with templates
	tmpl, err := getTemplate("shop.html")
	if err != nil {
		// TODO: fix nil logger bug (try delve?)
		// sm.logger.Error("template", (err))
		fmt.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ShoppingData{User: user, Products: products})
	if err != nil {
		// sm.logger.Error("template execute", (err))
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
