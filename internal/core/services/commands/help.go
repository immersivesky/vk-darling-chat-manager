package commands

import (
	"fmt"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/core/domain"
)

//nolint:all
var menuKeyboard = keyboard.Create().Inline().
	Callback("🛍️ Plugins", `{"command": "plugins"}`).
	Callback("✏ Control", `{"command": "control"}`).
	JSON()

type HelpCmd struct{}

func (c *HelpCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	bot.SendMessage(API.SendMessage{
		ChatID:   payload.Message.ChatID,
		Text:     fmt.Sprintf("%d, menu:", payload.Chat.ID),
		Keyboard: menuKeyboard,
		Forward:  API.GetForward(payload.Message.ChatID, payload.Message.ChatMessageID, true),
	})
}
