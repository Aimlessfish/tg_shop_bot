package controller

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Aimlessfish/tg_shop_bot/api"
	handler "github.com/Aimlessfish/tg_shop_bot/app/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

var lastMessageMap = make(map[int64]int)

func StartBot() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Shop")

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		logger.Warn("Bot token is missing or broken in the environment variables.")
		return fmt.Errorf("bot token is missing or broken")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Warn("Error running NewBot", "Error", err.Error())
		return err
	}

	logger.Info(fmt.Sprintf("Connected to account %v", bot.Self.UserName))

	db, err := api.DbInit()
	if err != nil {
		logger.Warn("Error", "Running api.dbInit failed: ", err.Error())
		os.Exit(1)
	}
	defer db.Close() // REMOVE

	//update handler
	update_channel := tgbotapi.NewUpdate(0)
	update_channel.Timeout = 60
	updates := bot.GetUpdatesChan(update_channel)

	for update := range updates {
		if update.Message != nil { //manage text
			logger.Info("Received message update", "chatID", update.Message.Chat.ID, "text", update.Message.Text)
			go HandleIncomingMessage(bot, update)
		} else if update.CallbackQuery != nil { //manage button presses
			logger.Info("Received callback query!", "callbackData", update.CallbackQuery.Data)
			go handler.HandleCallbackQuery(bot, update)
		}
	}
	return nil
}

func HandleIncomingMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	// Check if the message is not nil
	if update.Message != nil {
		// If it's a command, process it with CommandControl
		if update.Message.IsCommand() {
			CommandControl(bot, update.Message)
		} else if update.CallbackQuery != nil { // Handle callback query if present
			go handler.HandleCallbackQuery(bot, update) // Call HandleCallbackQuery function
		}
	}
	return nil // Return nil if no errors
}

func CommandControl(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "CommandControl")
	switch message.Command() {
	case "start":
		handler.HandleStart(bot, message)
	case "help":
		handler.HandleHelp(bot, message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command - Use /help for help!")
		if _, err := bot.Send(msg); err != nil {
			logger.Warn("Error handling unknown command msg", "error", err.Error())
			return err
		}
	}
	return nil
}
