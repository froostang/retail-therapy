package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/froostang/retail-therapy/api/user"
)

type totals struct {
	Total    string
	Tax      string
	Subtotal string
}

func (sm *ShoppingManager) getTotals() totals {
	result := totals{}
	cartContents := sm.cart.GetAll()

	var subtotal float64
	for _, item := range cartContents {
		s, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return result
		}
		subtotal += s
	}

	result.Subtotal = fmt.Sprintf("%.2f", subtotal)
	total := subtotal + subtotal*defaultTaxRate
	result.Total = fmt.Sprintf("%.2f", total)
	result.Tax = fmt.Sprintf("%.2f", total-subtotal)

	return result
}

func (sm *ShoppingManager) CheckoutRenderHandler(w http.ResponseWriter, r *http.Request) {
	t, err := getTemplate("checkout_updated.gohtml")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	totals := sm.getTotals()
	cartContents := sm.cart.GetAll()
	cartItemCount := len(cartContents)

	err = t.Execute(w, ShoppingData{
		User:          user.User{Name: "checkout man"},
		Products:      cartContents,
		Tax:           totals.Tax,
		Total:         totals.Total,
		Subtotal:      totals.Subtotal,
		CartItemCount: cartItemCount,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
