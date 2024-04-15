package utils

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//унеcти логгер в утилзы

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func ProvideLogger() fx.Option {
	return fx.Provide(NewLogger)
}
