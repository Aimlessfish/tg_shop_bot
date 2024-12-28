./app/checkout/
Here is the checkout handler for the shop.
This will handle button creation as well as assigning dynamicly enerated crypto payment addresses.

Example button creation module:

function to generate button text {
    buttons := [][]string{
        {"checkout", "back"},
    }
}
return controller.CreateInlineKeyboard(buttons)

github.com/go-telegram-bot-api/telegram-bot-api/v5