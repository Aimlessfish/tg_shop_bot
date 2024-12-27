package controller

import (
	"log/slog"
	"os"
)

func controller() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Shop")

	logger.Info("Hello")
}
