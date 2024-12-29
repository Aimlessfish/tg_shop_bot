package controller

import (
	"fmt"
	"log/slog"
	"os"

	index "github.com/Aimlessfish/tg_shop_bot/index"
	shop "github.com/Aimlessfish/tg_shop_bot/shop"
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
	chatID := message.Chat.ID
	keyboard := index.Buttons()

	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome! Please use the buttons to navigate the store.")
	msg.ReplyMarkup = keyboard

	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error", "Failed to send keyboard", err.Error())
		os.Exit(1)
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please use /start to start!")
	bot.Send(msg)
	return nil
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleCallbackQuery")
	query := update.CallbackQuery
	chatID := query.Message.Chat.ID

	if lastMsgID, exists := lastMessageMap[chatID]; exists {
		deleteConfig := tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: lastMsgID,
		}
		_, err := bot.Request(deleteConfig)
		if err != nil {
			logger.Warn("Error deleting previous message", "Error: ", err.Error())
		}
		logger.Info("Passed previous message check!")
	}

	//log the callback data & respond
	logger.Info("Callback received!", "Data: ", query.Data)
	response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
	_, err := bot.Request(response)
	if err != nil {
		logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
		os.Exit(1)
	}

	//handle the callback data
	var messageText string
	var keyboard tgbotapi.InlineKeyboardMarkup

	switch query.Data {
	case "shop":
		messageText = "Displaying all items in shop!"
		keyboard = shop.Buttons()
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, messageText)
		msg.ReplyMarkup = keyboard
		sentMsg, err := bot.Send(msg)
		if err != nil {
			logger.Warn("Error", "Failed to follow up the callback query", err.Error())
			return err
		}
		lastMessageMap[chatID] = sentMsg.MessageID

	case "support":
		if lastMsgID, exists := lastMessageMap[chatID]; exists {
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: lastMsgID,
			}
			_, err := bot.Request(deleteConfig)
			if err != nil {
				logger.Warn("Error deleting previous message", "Error: ", err.Error())
			}
			logger.Info("Passed previous message check!")
		}
		messageText = "Please contact @username for support"
		keyboard = index.Buttons()
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, messageText)
		msg.ReplyMarkup = keyboard
		sentMsg, err := bot.Send(msg)
		if err != nil {
			logger.Warn("Error", "Failed to follow up the callback query", err.Error())
			return err
		}
		lastMessageMap[chatID] = sentMsg.MessageID

	case "tracking":
		messageText = "Please wait while we get your tracking number..."

	case "orders":
		if lastMsgID, exists := lastMessageMap[chatID]; exists {
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: lastMsgID,
			}
			_, err := bot.Request(deleteConfig)
			if err != nil {
				logger.Warn("Error deleting previous message", "Error: ", err.Error())
			}
			logger.Info("Passed previous message check!")
		}

		messageText = "Please wait while we retrieve your previous orders..."
		keyboard = index.Buttons()
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, messageText)
		msg.ReplyMarkup = keyboard
		sentMsg, err := bot.Send(msg)
		if err != nil {
			logger.Warn("Error", "Failed to follow up the callback query", err.Error())
			return err
		}
		lastMessageMap[chatID] = sentMsg.MessageID

	case "back":
		if lastMsgID, exists := lastMessageMap[chatID]; exists {
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: lastMsgID,
			}
			_, err := bot.Request(deleteConfig)
			if err != nil {
				logger.Warn("Error deleting previous message", "Error: ", err.Error())
			}
			logger.Info("Passed previous message check!")
		}

		response := tgbotapi.NewCallback(query.ID, "Taking you back...")
		if _, err := bot.Request(response); err != nil {
			logger.Warn("Error sending callback response", "error", err.Error())
			return err
		}
		keyboard := index.Buttons()
		msg := tgbotapi.NewMessage(chatID, "Taking you back...")
		msg.ReplyMarkup = keyboard

		sentMsg, err := bot.Send(msg)
		if err != nil {
			logger.Warn("Error sending new message with buttons", "error", err.Error())
			return err
		}
		lastMessageMap[chatID] = sentMsg.MessageID
	}

	return nil
}
