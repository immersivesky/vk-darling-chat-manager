package commands

import (
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
)

//nolint:all
var menuKeyboard = keyboard.Create().Inline().
	Callback("🛍️ Plugins", `{"command": "plugins"}`).
	Callback("✏ Control", `{"command": "control"}`).
	JSON()

func helpScript(bot *API.Bot, payload *Payload) {
	bot.SendMessage(API.SendMessage{
		ChatID:   payload.message.ChatID,
		Text:     "UNSPECIFIED, menu:",
		Keyboard: menuKeyboard,
		Forward:  API.GetForward(payload.message.ChatID, payload.message.ChatMessageID, true),
	})
}
