package shop

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Catergories() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "shopButtons: catergories")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Hats", "category"),
			tgbotapi.NewInlineKeyboardButtonData("Coats", "category"),
			tgbotapi.NewInlineKeyboardButtonData("Main Menu", "back_main"),
		},
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}

func Listings() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "shopButtons: Listings")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Hat", "item"),
			tgbotapi.NewInlineKeyboardButtonData("Previous Page", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Main Menu", "back_main"),
		},
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}
