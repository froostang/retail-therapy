package loggers

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (zl *ZapLogger) Error(msg string, err error) {
	zl.logger.Error(msg, zap.Error(err))
}

func (zl *ZapLogger) Info(msg ...string) {
	// FIXME: This is a hack and not idea long term
	for _, m := range msg {
		zl.logger.Info(m)
	}
}
