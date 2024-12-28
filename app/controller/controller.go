package controller

import (
	"fmt"
	"log/slog"
	"os"

	index "github.com/Aimlessfish/tg_shop_bot/index"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

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
		logger.Warn("Bot token is missing in the environment variables.")
		return fmt.Errorf("bot token is missing")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Warn("Error running NewBot", "Error", err.Error())
		return err
	}

	logger.Info(fmt.Sprintf("Connected to account %v", bot.Self.UserName))

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
			go HandleCallbackQuery(bot, update)
		}
	}
	return nil
}

func HandleIncomingMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	// Check if the message is not nil
	if update.Message != nil {
		// If it's a command, process it with CommandControl
		if update.Message.IsCommand() {
			CommandControl(bot, update.Message) // Corrected: use `update.Message`
		} else if update.CallbackQuery != nil { // Handle callback query if present
			HandleCallbackQuery(bot, update) // Call HandleCallbackQuery function
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
		HandleStart(bot, message)
	case "help":
		HandleHelp(bot, message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command - Use /help for help!")
		if _, err := bot.Send(msg); err != nil {
			logger.Warn("Error handling unknown command msg", "error", err.Error())
			return err
		}
	}
	return nil
}

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleStart")

	keyboard := index.Buttons()

	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome! Please use the buttons to navigate the store.")
	msg.ReplyMarkup = keyboard

	_, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error", "Failed to send keyboard", err.Error())
		os.Exit(1)
	}

	return nil
}

func HandleHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please use /start to start!")
	bot.Send(msg)
	return nil
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	query := update.CallbackQuery
	response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
	bot.Request(response)
	return nil
}

// func GenerateInlineKeyboard(buttons [][]string) (tgbotapi.InlineKeyboardMarkup, error) {
// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
// 	slog.SetDefault(logger)
// 	logger = slog.With("LogID", "InlineGenerator")

// 	// Create an empty slice for Telegram inline keyboard rows
// 	var keyboardRows [][]tgbotapi.InlineKeyboardButton

// 	// Iterate over the input to construct the inline keyboard rows
// 	for _, row := range buttons {
// 		var keyboardRow []tgbotapi.InlineKeyboardButton
// 		for _, buttonText := range row {
// 			// Create a button with callback data (can customize callback data as needed)
// 			button := tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonText)
// 			keyboardRow = append(keyboardRow, button)
// 		}
// 		keyboardRows = append(keyboardRows, keyboardRow)
// 	}

// 	// Create the final inline keyboard markup
// 	inlineKeyboard := tgbotapi.InlineKeyboardMarkup{
// 		InlineKeyboard: keyboardRows,
// 	}

// 	logger.Info("Generated inline keyboard successfully")
// 	return inlineKeyboard, nil
// }
