package index

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Buttons() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "buttons: index")

	buttons := [][]tgbotapi.InlineKeyboardButton{

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Shop", "shop"),
			tgbotapi.NewInlineKeyboardButtonData("Support", "support"),
			tgbotapi.NewInlineKeyboardButtonData("Tracking", "tracking"),
			tgbotapi.NewInlineKeyboardButtonData("Orders", "orders"),
		),
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}
