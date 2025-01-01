package handler

import (
	"fmt"
	"log/slog"
	"os"

	index "github.com/Aimlessfish/tg_shop_bot/app/index"
	orders "github.com/Aimlessfish/tg_shop_bot/app/previous"
	shop "github.com/Aimlessfish/tg_shop_bot/app/shop"
	tracking "github.com/Aimlessfish/tg_shop_bot/app/tracking"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var lastMessageMap = make(map[int64]int)

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "HandleStart")
	chatID := message.Chat.ID
	keyboard := index.Buttons()
	var username string

	//HANDLE DB CONNECT
	//RETURN VENDOR_ID, SALES, RATING, LAST_SEEN

	username = message.From.UserName
	msg := tgbotapi.NewMessage(message.Chat.ID, " { .VENDOR_NAME }\nSales:{ .VENDOR_SALES }\nRating:{ .VENDOR_RATING }\nLast Seen:{ .LAST_SEEN_STATUS }")
	bot.Send(msg)
	msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Welcome <b>%v</b>! Please use the buttons to navigate the store", username))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	sentMsg, err := bot.Send(msg)
	if err != nil {
		logger.Warn("Error", "Failed to send keyboard", err.Error())
		os.Exit(1)
	}
	lastMessageMap[chatID] = sentMsg.MessageID

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

	lastMessageMap[chatID] = query.Message.MessageID

	switch query.Data {
	case "shop":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleShop(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleShop", err.Error())
			return err
		}

	case "support":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleSupport(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleSupport", err.Error())
		}

	case "tracking":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleTracking(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleTracking", err.Error())
			return err
		}

	case "orders":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandlePreviousOrders(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandlePreviousOrders", err.Error())
			return err
		}

	case "category":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleListings(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleListings", err.Error())
			return err
		}
	case "item":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleItem(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleItem", err.Error())
			return err
		}
	case "back":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleBackButton(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandlePreviousOrders", err.Error())
			return err
		}

	case "back_main":
		logger.Info("Callback received!", "Data: ", query.Data)
		response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
		_, err := bot.Request(response)
		if err != nil {
			logger.Warn("Error sending callback response for query: ", query.Data, err.Error())
			os.Exit(1)
		}
		err = HandleMainMenu(bot, query.Message)
		if err != nil {
			logger.Warn("Error: ", "Callback query failed. Case: HandleShop", err.Error())
			return err
		}

	}
	return nil
}

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
