package commands

import (
	"fmt"
	"github.com/botscommunity/vkgo/API"
	"time"
)

var pingPattern = "^(?:ping)$"

func pingScript(bot *API.Bot, payload *Payload) {
	var (
		text = fmt.Sprintf("%s, pong!", "UNSPECIFIED")
		sent = bot.SendMessage(payload.message.ChatID, text)
	)

	if sent.Error.Message == "" {
		bot.EditMessage(
			sent.ChatID,
			sent.ChatMessageID,
			fmt.Sprintf("%s\nDuration: %d ms.",
				text,
				int(time.Now().UnixMilli())-payload.message.Date,
			),
		)
	}
}
