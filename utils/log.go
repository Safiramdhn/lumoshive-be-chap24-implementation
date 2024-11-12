package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func IntiLogger() *zap.Logger {
	file, _ := os.Create("app.log")
	writer := zapcore.AddSync(file)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		writer,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	return logger
}
