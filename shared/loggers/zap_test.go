package loggers_test

import (
	"fmt"
	"testing"

	"github.com/froostang/retail-therapy/shared/loggers"
	"go.uber.org/zap/zaptest"
)

func TestZap(t *testing.T) {
	logger := zaptest.NewLogger(t)

	l := loggers.NewZapLogger(logger)

	l.Error("hello", fmt.Errorf("fakeErr"))

}
