package controller

import (
	"log/slog"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)

func StartBot() err{
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger = slog.With("LogID", "Shop")

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
	return nil

	//update handler
	update_channel := tgbotapi.NewUpdate(0)
	update_channel.Timeout = 60 
	updates := bot.GetUpdatesChan(update_channel)

	for update := range updates {
		if update.Message != nil { //manage text	
			logger.Info("Received message update", "chatID", update.Messahe.Chat.ID, "text", update.Message.text)
			go HandleIncomingMessage(bot, update)
		} else if update.CallbackQuery != nil { //manage button presses
			logger.Info("Received callback query!", "callbackData", update.CallbackQuery.Data)
			go HandleCallbackQuery(bot, update)
		}
	}
	return nil
}

func HandleIncomingMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) err {
	if update.Message != nil {
		if update.Message.IsCommand() {
			CommandControl(bot, update.Messages)
		} else if update.CallbackQuery != nil {
			HandleCallbackQuery(bot, update)
		}
		}
	}
	return nil
}

func CommandControl(bot *tgbotapi.BotAPI, message *tgbotapi.Message) err {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
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

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) err {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome! Please use the buttons to navigate the store.")
	bot.Send(msg)
	return nil
}

func HandleHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) err {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please use /start to start!")
	bot.Send(msg)
	return nil
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) err {
	query := update.CallbackQuery
	response := tgbotapi.NewCallback(query.ID, fmt.Sprintf("Taking you to %v", query.Data))
	bot.Request(response)
	return nil
}


func GenerateInlineKeyboard() err{
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = slog.With("LogID", "InlineGenerator")


}
