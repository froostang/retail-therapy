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

	// Apply JWT middleware to all routes
	http.Handle("/shop", middleware.Apply(http.HandleFunc("/shop", handlers.NewShoppingManager(handlers.AddLogger(loggers.NewZapLogger(logger))).ShoppingHandler),
		middleware.PanicRecovery, middleware.JWTMiddleware))
	// Apply panic recover middleware

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
