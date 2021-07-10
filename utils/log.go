package utils

import (
	"log"

	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	z, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v\n", err)
	}
	return z
}
