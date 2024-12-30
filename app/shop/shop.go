package shop

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// fun Categories() will handle api connection for database queries specific to vender_categories
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

// fun Listing() will handle api connection for database queries specific to vender_listings
func Listings() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "shopButtons: Listings")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Hat", "item")},
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("Back", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Main Menu", "back_main"),
		),
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}

func Item() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "shopButtons: Items")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Quantity +", "quantity+"),
			tgbotapi.NewInlineKeyboardButtonData(" {.value}", "quantity"),
			tgbotapi.NewInlineKeyboardButtonData("Quantity -", "quantity-"),
		},
		{tgbotapi.NewInlineKeyboardButtonData("Add to Basket", "basket_add")},
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Back", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Main Menu", "back_main"),
		),
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}
