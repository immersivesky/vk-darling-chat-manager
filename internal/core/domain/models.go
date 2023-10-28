package domain

import (
	"time"

	"github.com/botscommunity/vkgo/update"
)

type Chat struct {
	ChatID int
}

type ChatMember struct {
	Name string
}

type Payload struct {
	Message    update.Message
	Chat       *Chat
	ChatMember *ChatMember
	Time       time.Time
}
