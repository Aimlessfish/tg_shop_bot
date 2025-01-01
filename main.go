package main

import (
	"log/slog"
	"os"

	controller "github.com/Aimlessfish/tg_shop_bot/app/controller"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Main")

	db, err := api.dbInit()
	if err != nil {
		logger.Warn("Error", "Running api.dbInit failed: ", err.Error())
		os.Exit(1)
	}

	err = controller.StartBot()
	if err != nil {
		logger.Warn("Error starting shop!", "controller.StartBot", err.Error())
		os.Exit(1)
	}
}
