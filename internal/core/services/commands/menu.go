package commands

import (
	"fmt"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
)

//nolint:all
var menuKeyboard = keyboard.Create().Inline().
	Text("🛍️ Plugins", `{"command": "plugins"}`).
	Text("✏ Control", `{"command": "control"}`).
	JSON()

type MenuCmd struct{}

func (c *MenuCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	bot.SendMessage(API.SendMessage{
		ChatID:   payload.Message.ChatID,
		Text:     fmt.Sprintf("%s, menu:", payload.ChatMember.Name),
		Keyboard: menuKeyboard,
		Forward:  API.GetForward(payload.Message.ChatID, payload.Message.ChatMessageID, true),
	})
}
