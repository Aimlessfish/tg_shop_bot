package handler

import (
	"log/slog"
	"os"

	index "github.com/Aimlessfish/tg_shop_bot/app/index"
	orders "github.com/Aimlessfish/tg_shop_bot/app/previous"
	shop "github.com/Aimlessfish/tg_shop_bot/app/shop"
	tracking "github.com/Aimlessfish/tg_shop_bot/app/tracking"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var lastMessageMap = make(map[int64]int)

func HandleHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please use /start to start!")
	bot.Send(msg)
	return nil
}

func HandleShop(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleShop")

	chatID := message.Chat.ID
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

	messageText := "Displaying all categories in shop!"
	keyboard := shop.Catergories()
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	msg.ReplyMarkup = keyboard
	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error: ", "Failed serve func HandleShop", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleSupport(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleShop")

	chatID := message.Chat.ID
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
	messageText := "Please contact @username for support"
	keyboard := index.Buttons()
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	msg.ReplyMarkup = keyboard
	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error", "Failed to follow up the callback query", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandlePreviousOrders(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleShop")

	chatID := message.Chat.ID
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

	messageText := "Please wait while we retrieve your previous orders..."
	keyboard := orders.Buttons()
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	msg.ReplyMarkup = keyboard
	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error", "Failed to follow up the callback query", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleTracking(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleTracking")

	chatID := message.Chat.ID
	messageText := "Please wait while we get your tracking number..."

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
	keyboard := tracking.Buttons()
	msg := tgbotapi.NewMessage(chatID, messageText)
	msg.ReplyMarkup = keyboard
	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error sending new message with buttons", "error", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleListings(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleListings")

	chatID := message.Chat.ID

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
	keyboard := shop.Listings()
	msg := tgbotapi.NewMessage(chatID, "Listing all items!")
	msg.ReplyMarkup = keyboard

	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error sending new message with buttons", "error", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleItem(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleItem")

	chatID := message.Chat.ID

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
	keyboard := shop.Item()
	msg := tgbotapi.NewMessage(chatID, "{ .ItemName }\n{ .ItemDescription }\n{ .Prices }")
	msg.ReplyMarkup = keyboard

	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error sending new message with buttons", "error", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}

func HandleBackButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleMainMenu")

	chatID := message.Chat.ID

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
	return nil
}

func HandleMainMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleMainMenu")

	chatID := message.Chat.ID

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
	keyboard := index.Buttons()
	msg := tgbotapi.NewMessage(chatID, "Main Menu")
	msg.ReplyMarkup = keyboard

	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error sending new message with buttons", "error", err.Error())
		return err
	}
	lastMessageMap[chatID] = sentMsg.MessageID

	return nil
}
