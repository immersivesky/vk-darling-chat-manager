package scenes

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/update"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/ports"
	"gitlab.com/immersivesky/affinitycm-vk/pkg/consts"
)

type MessageSceneService struct {
	db              ports.DB
	publicCommands  ports.Commands
	privateCommands ports.Commands
}

func NewMessageSceneService(db ports.DB, publicCommandsSvc ports.Commands, privateCommandsSvc ports.Commands) *MessageSceneService {
	return &MessageSceneService{
		db:              db,
		publicCommands:  publicCommandsSvc,
		privateCommands: privateCommandsSvc,
	}
}

func (ms *MessageSceneService) GetMessageScene(bot *API.Bot, object update.Object) {
	payload := &domain.Payload{
		Message: object.Message,
		Time:    time.Now(),
	}

	chat, err := ms.db.GetChat(payload.Message.ChatID)
	payload.Chat = chat

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			payload.Chat, err = ms.db.CreateChat(payload.Message.ChatID)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	if object.Message.ChatID > consts.ChatsStartIn {
		for regex, cmd := range ms.publicCommands {
			if ok := regex.MatchString(payload.Message.Text); ok {
				cmd.Execute(bot, payload)
			}
		}
	} else {
		for regex, cmd := range ms.privateCommands {
			if ok := regex.MatchString(payload.Message.Text); ok {
				cmd.Execute(bot, payload)
			}
		}
	}
}
