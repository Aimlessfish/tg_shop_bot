package main

import (
	"fmt"
	"log/slog"
	"os"
	controller "./controller/controller"

)

func main() err{
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Main")

	err := controller.StartBot()
	if err != nil {
		logger.Warn("Error starting shop!", "controller.StartBot", err.Error())
		os.Exit(1)
	}
	return nil
}
