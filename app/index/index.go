package index

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Buttons() (tgbotapi.InlineKeyboardMarkup, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "indexButtons")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Shop", "shop"),
			tgbotapi.NewInlineKeyboardButtonData("Support", "support"),
			tgbotapi.NewInlineKeyboardButtonData("Tracking", "tracking"),
		},
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	return keyboard, nil
}
