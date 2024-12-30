package controller

import (
	"fmt"
	"log/slog"
	"os"

	handler "github.com/Aimlessfish/tg_shop_bot/handlers"
	index "github.com/Aimlessfish/tg_shop_bot/index"
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
			CommandControl(bot, update.Message)
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
		err = handler.HandleShop(bot, query.Message)
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
		err = handler.HandleSupport(bot, query.Message)
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
		err = handler.HandleTracking(bot, query.Message)
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
		err = handler.HandlePreviousOrders(bot, query.Message)
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
		err = handler.HandleListings(bot, query.Message)
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
		err = handler.HandleItem(bot, query.Message)
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
		err = handler.HandleBackButton(bot, query.Message)
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
