package http

import (
	"fmt"
	"net/http"

	"github.com/froostang/retail-therapy/api/http/handlers"
	"github.com/froostang/retail-therapy/shared/loggers"
	"github.com/froostang/retail-therapy/shared/middleware"
	"go.uber.org/zap"
)

func NewServer(logger *zap.Logger) {

	// TODO: Apply JWT middleware to all routes middleware.JWTMiddleware
	logger.Info("registering handlers")

	// TODO: just let it inline it
	shoppingManager := &handlers.ShoppingManager{}
	shoppingManager = handlers.NewShoppingManager(shoppingManager,
		handlers.AddLogger(loggers.NewZapLogger(logger)))

	http.Handle("/shop", middleware.Apply(
		http.HandlerFunc(shoppingManager.ShoppingHandler),
		middleware.PanicRecovery))

	AdderManager := handlers.NewAdderManager(loggers.NewZapLogger(logger))
	http.Handle("/add-product", middleware.Apply(
		http.HandlerFunc(AdderManager.AdderHandler),
		middleware.PanicRecovery))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
