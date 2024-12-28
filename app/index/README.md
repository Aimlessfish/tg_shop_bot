./app/index/
Containing the index or start menu of the shop
and serve the layout  which provides the contents through the API layer which will call the backend to pull from the db: for now we have basic plain text buttons to serve development use.

buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Shop", "shop"),
			tgbotapi.NewInlineKeyboardButtonData("Support", "support"),
			tgbotapi.NewInlineKeyboardButtonData("Tracking", "tracking"),
		},
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard