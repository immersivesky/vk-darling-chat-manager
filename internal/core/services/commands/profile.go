package commands

import (
	"fmt"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"

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
		Text: fmt.Sprintf(`📒 %s, your profile:

🆔 ID: %d
💎 Balance: 🍓 0 🥑 0 ☕ 0`, payload.ChatMember.Name, payload.Message.UserID),
		Keyboard: profileKeyboard,
		Forward:  API.GetForward(payload.Message.ChatID, payload.Message.ChatMessageID, true),
	})
}
