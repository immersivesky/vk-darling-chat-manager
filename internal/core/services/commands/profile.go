package commands

import (
	"fmt"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/core/domain"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
)

//nolint:all
var profileKeyboard = keyboard.Create().Inline().
	Callback("📅 Notes").JSON()

type ProfileCmd struct{}

func (c *ProfileCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	bot.SendMessage(API.SendMessage{
		ChatID: payload.Message.ChatID,
		Text: fmt.Sprintf(`📒 %d, your profile:

🆔 ID: 10
💎 Balance: 🍓 0 🥑 0 ☕ 0`, payload.Chat.ID),
		Keyboard: profileKeyboard,
		Forward:  API.GetForward(payload.Message.ChatID, payload.Message.ChatMessageID, true),
	})
}
