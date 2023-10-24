package commands

import (
	"fmt"
	"time"

	"github.com/botscommunity/vkgo/API"
)

func pingScript(bot *API.Bot, payload *Payload) {
	var (
		text = fmt.Sprintf("%s, pong!", payload.user.name)
		sent = bot.SendMessage(payload.message.ChatID, text)
	)

	if sent.Error.Message == "" {
		bot.EditMessage(
			sent.ChatID,
			sent.ChatMessageID,
			fmt.Sprintf("%s\nDuration: %d",
				text,
				time.Since(payload.Time),
			),
		)
	}
}
