package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

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

	// stockImg := "https://images.unsplash.com/photo-1591279304068-c997c097f2b7?q=80&w=2970&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"

	// products := []product.Scraped{{Name: "test1", ImageURL: stockImg}, {Name: "test2", ImageURL: stockImg}}

	products := sm.cache.Get()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	// TODO: fix directory structure issues with templates
	tmpl, err := template.ParseFiles(exPath + "/templates/shop.html")
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
