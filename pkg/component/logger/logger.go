package logger

import (
	"github.com/solate/util/logger"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init(level, path string) (err error) {
	Logger, err = logger.NewLogger(level, path)
	return
}
