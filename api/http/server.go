package http

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/froostang/retail-therapy/api/cache"
	"github.com/froostang/retail-therapy/api/http/handlers"
	"github.com/froostang/retail-therapy/api/shared/loggers"
	"github.com/froostang/retail-therapy/api/shared/middleware"
	"go.uber.org/zap"
)

var (
	globalProductCache *cache.Products
	globalCart         *cache.Products
	TemplateFS         embed.FS
)

func NewServer(logger *zap.Logger, templates embed.FS) {
	handlers.SetTemplates(templates)
	// TODO: Apply JWT middleware to all routes middleware.JWTMiddleware
	logger.Info("registering handlers")
	globalProductCache = cache.NewForProducts(loggers.NewZapLogger(logger), 100)
	globalCart = cache.NewForProducts(loggers.NewZapLogger(logger), 100)

	shoppingManager := &handlers.ShoppingManager{}
	shoppingManager = handlers.NewShoppingManager(shoppingManager,
		handlers.AddLogger(loggers.NewZapLogger(logger)),
		handlers.AddCacher(globalProductCache),
		handlers.AddCart(globalCart))

	// Shop view
	http.Handle("/shop", middleware.Apply(
		http.HandlerFunc(shoppingManager.ShoppingHandler),
		middleware.PanicRecovery))

	// Cart add
	http.Handle("/add-cart", middleware.Apply(
		http.HandlerFunc(shoppingManager.CartHandler),
		middleware.PanicRecovery))

	http.Handle("/checkout", middleware.Apply(
		http.HandlerFunc(shoppingManager.CheckoutRenderHandler),
		middleware.PanicRecovery))

	http.Handle("/complete", middleware.Apply(
		http.HandlerFunc(shoppingManager.CompleteHandler),
		middleware.PanicRecovery))

	// Add POST
	AdderManager := handlers.NewAdderManager(loggers.NewZapLogger(logger), globalProductCache)
	http.Handle("/add", middleware.Apply(
		http.HandlerFunc(AdderManager.AdderHandler),
		middleware.PanicRecovery))

	// Add view
	http.Handle("/about", middleware.Apply(
		http.HandlerFunc(handlers.AboutRenderHandler),
		middleware.PanicRecovery))

	// Add view
	http.Handle("/add-product", middleware.Apply(
		http.HandlerFunc(handlers.AdderRenderHandler),
		middleware.PanicRecovery))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
