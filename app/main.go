package main

import (
	"log/slog"
	"os"

	controller "github.com/Aimlessfish/tg_shop_bot/controller"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Main")

	err := controller.StartBot()
	if err != nil {
		logger.Warn("Error starting shop!", "controller.StartBot", err.Error())
		os.Exit(1)
	}
}
