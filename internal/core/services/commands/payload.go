package commands

import "github.com/botscommunity/vkgo/update"

type Payload struct {
	message update.Message
	user    struct {
		name string
	}
}

func NewCommandPayload(message update.Message) *Payload {
	return &Payload{
		message: message,
	}
}
