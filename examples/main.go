package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgbot/conversation"
)

var conv = conversation.New()

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	message, err := ctx.EffectiveMessage.Reply(
		b,
		"Do you like to tell me your name?",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
					{
						{
							Text:         "Yes",
							CallbackData: "yes",
						},
					},
					{
						{
							Text:         "No",
							CallbackData: "no",
						},
					},
				},
			},
		},
	)
	if err != nil {
		return nil
	}

	update := conv.WaitForCallback(ctx, message)
	if update == nil {
		return nil
	}

	callback := update.CallbackQuery.Data
	if callback == "no" {
		message.EditText(b, "OK, have a great time!", nil)
		return nil
	}

	message.EditText(b, "Cool!", nil)
	ctx.EffectiveMessage.Reply(b, "So what's your first name?", nil)
	update = conv.WaitForMessage(ctx, conversation.Text)
	if update == nil {
		return nil
	}

	name := update.Message.Text
	ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Oh, hi %s!", name), nil)
	return nil
}

func main() {
	bot, err := gotgbot.NewBot("1925137673:AAEh2-9MzbS-ISOWtMvzoN3ITZwpQ0ExO4w", nil)
	if err != nil {
		panic(err)
	}

	updater := ext.NewUpdater(nil)
	updater.Dispatcher.AddHandlerToGroup(conv, -1)
	updater.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	updater.StartPolling(bot, &ext.PollingOpts{DropPendingUpdates: true})
	updater.Idle()
}
