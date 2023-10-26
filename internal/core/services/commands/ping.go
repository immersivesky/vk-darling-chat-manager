package commands

import (
	"fmt"
	"time"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/core/domain"

	"github.com/botscommunity/vkgo/API"
)

type PingCmd struct{}

func (c *PingCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	var (
		beforeMessageDuration = time.Since(payload.Time)
		text                  = fmt.Sprintf("%d, pong!", payload.Chat.ID)
		sent                  = bot.SendMessage(payload.Message.ChatID, text)
	)

	if sent.Error.Message == "" {
		bot.EditMessage(
			sent.ChatID,
			sent.ChatMessageID,
			fmt.Sprintf("%s\nDuration: %v\nWithout SendMessage: %v",
				text,
				time.Since(payload.Time),
				beforeMessageDuration,
			),
		)
	}
}
