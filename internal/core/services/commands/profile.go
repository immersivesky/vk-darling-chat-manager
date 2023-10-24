package commands

import (
	"fmt"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
)

//nolint:all
var profileKeyboard = keyboard.Create().Inline().
	Callback("📅 Notes").JSON()

func profileScript(bot *API.Bot, payload *Payload) {
	bot.SendMessage(API.SendMessage{
		ChatID: payload.message.ChatID,
		Text: fmt.Sprintf(`📒 %s, your profile:

🆔 ID: 10
💎 Balance: 🍓 0 🥑 0 ☕ 0`, payload.message.Text),
		Keyboard: profileKeyboard,
		Forward:  API.GetForward(payload.message.ChatID, payload.message.ChatMessageID, true),
	})
}
