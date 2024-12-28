package index

import (
	"log/slog"
	"os"

	controller "github.com/Aimlessfish/tg_shop_bot/controller"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func indexButtons() (tgbotapi.InlineKeyboardMarkup, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "indexButtons")

	buttons := [][]string{
		{"Shop", "Support", "Tracking"},
	}
	return controller.GenerateInlineKeyboard(buttons)
}
