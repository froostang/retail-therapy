package handlers

import (
	"html/template"
	"net/http"

	"github.com/froostang/retail-therapy/api/product"
	"github.com/froostang/retail-therapy/api/user"
)

type ShoppingData struct {
	User     user.User
	Products []product.Scraped
}

func ShoppingHandler(w http.ResponseWriter, r *http.Request) {
	user := user.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	products := []product.Scraped{{Name: "test1", ImageURL: "https://images.unsplash.com/photo-1591279304068-c997c097f2b7?q=80&w=2970&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"}}

	tmpl, err := template.ParseFiles("templates/shopping.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ShoppingData{User: user, Products: products})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
