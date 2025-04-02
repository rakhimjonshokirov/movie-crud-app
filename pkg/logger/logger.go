package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"logger",
	fx.Provide(NewLogger),
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
