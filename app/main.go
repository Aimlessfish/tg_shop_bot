package main

import (
	"fmt"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Main")

	bot, err := tgbotapi.NewBotAPI("") // call from .env
	if err != nil {
		logger.Warn("Error running NewBot", "Error", err.Error())
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("Connected to account %v", bot.Self.UserName))
}
