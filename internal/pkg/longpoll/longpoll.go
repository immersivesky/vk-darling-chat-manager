package longpoll

import (
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/longpoll"
	"github.com/botscommunity/vkgo/scene"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/scenes"
)

type Longpoll struct {
	client *longpoll.Options
}

func NewLongpoll(
	token string,
	messageSceneService *scenes.MessageSceneService,
) *Longpoll {
	bot := API.Create(token)

	scenes := scene.Create().
		Message(messageSceneService.GetMessageScene)

	return &Longpoll{longpoll.Create(bot, scenes)}
}

func (s *Longpoll) Run() error {
	return s.client.Listen()
}
