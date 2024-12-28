./app/menu_index/
Containing the index or start menu of the shop
This calls controller.GenerateInlineKeyboard() to serve the layout and then provides the contents through the API layer which will call the backend to pull from the db: for now we have basic plain text buttons to serve development use.

indexButtons{
    buttons := [][]string{
        {"Shop", "Support", "Tracking"},
    }
}
return controller.CreateInlineKeyboard(buttons)