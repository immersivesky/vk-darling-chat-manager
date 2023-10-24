package commands

import (
	"time"

	"github.com/botscommunity/vkgo/update"
)

type Payload struct {
	message update.Message
	user    struct {
		name string
	}
	Time time.Time
}

func NewCommandPayload(message update.Message) *Payload {
	return &Payload{
		message: message,
	}
}
