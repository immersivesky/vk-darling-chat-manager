package commands

import (
	"fmt"
	"github.com/botscommunity/vkgo/API"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
)

type ControlCmd struct{}

func (c *ControlCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	bot.SendMessage(payload.Message.ChatID, fmt.Sprintf("%s, control menu:", payload.ChatMember.Name))
}
