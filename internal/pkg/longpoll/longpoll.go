package longpoll

import (
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/longpoll"
	"github.com/botscommunity/vkgo/scene"
)

type Longpoll struct {
	client *longpoll.Options
}

func NewLongpoll(
	token string,
	messageService scene.Message,
) *Longpoll {
	bot := API.Create(token)

	scenes := scene.Create().
		Message(messageService)

	return &Longpoll{longpoll.Create(bot, scenes)}
}

func (s *Longpoll) Run() error {
	return s.client.Listen()
}
