/app/controller/
This will handle the inline keyboard creation for reusable structures used by each module.

The controller will also route commands to their appropriate command handlers. 
Example:

func HandleIncomingMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	if update.Message != nil {
		if update.Message.IsCommand() { // If message is a command, process it with CommandControl
			CommandControl(bot, update.Message) 
		} else if update.CallbackQuery != nil { // Handle callback query if present
			HandleCallbackQuery(bot, update) // Call HandleCallbackQuery function
		}
	}
	return nil 
}

github.com/go-telegram-bot-api/telegram-bot-api/v5
