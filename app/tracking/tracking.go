package tracking

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Buttons() tgbotapi.InlineKeyboardMarkup {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "indexButtons")

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Main Menu", "back_main"),
		},
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}