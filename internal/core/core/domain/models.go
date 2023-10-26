package domain

import (
	"time"

	"github.com/botscommunity/vkgo/update"
)

type Chat struct {
	ID int
}

type Payload struct {
	Message update.Message
	Chat    *Chat
	Time    time.Time
}
