package handlers_test

import (
	"fmt"
	"testing"

	"github.com/froostang/retail-therapy/api/http/handlers"
	"github.com/froostang/retail-therapy/shared/loggers"
	"go.uber.org/zap/zaptest"
)

func TestManagerHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)

	l := loggers.NewZapLogger(logger)

	sm := &handlers.ShoppingManager{}
	handlers.NewShoppingManager(sm, handlers.AddLogger(l))

	l.Error("hello", fmt.Errorf("fakeErr"))
}
