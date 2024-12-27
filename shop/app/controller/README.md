/app/controller/
This will handle the inline keyboard creation for reusable structures used by each module.

Example:

function to generate button text {
    buttons := [][]string{
        {"add to basket", "back"},
    }
}
return shop.CreateInlineKeyboard(buttons)

github.com/go-telegram-bot-api/telegram-bot-api/v5
