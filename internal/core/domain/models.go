package domain

import (
	"github.com/botscommunity/vkgo/update"
	"time"
)

type Payload struct {
	Message    update.Message
	Chat       *Chat
	ChatMember *ChatMember
	Time       time.Time
}

type Chat struct {
	ChatID int
}

type ChatMember struct {
	Name string
}

type ChatPlugin struct {
	PluginID int
}

type Plugin struct {
	Status   int8
	Name     string
	Endpoint string
}
