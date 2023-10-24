package scenes

import (
	"time"

	"gitlab.com/immersivesky/affinitycm-vk/internal/pkg/consts"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/update"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
)

type MessageSceneService struct {
	commands *commands.Service
}

func NewMessageSceneService(commandsSvc *commands.Service) *MessageSceneService {
	return &MessageSceneService{
		commands: commandsSvc,
	}
}

func (ms *MessageSceneService) GetMessageScene(bot *API.Bot, object update.Object) {
	payload := commands.NewCommandPayload(object.Message)
	payload.Time = time.Now()

	if object.Message.ChatID > consts.ChatsStartIn {
		ms.commands.PublicExecute(bot, payload)
	} else {
		ms.commands.PrivateExecute(bot, payload)
	}
}
