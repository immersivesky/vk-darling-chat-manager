package scenes

import (
	"errors"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
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

	if _, ok := payload.Message.Payload["command"]; ok {
		payload.Message.Text = payload.Message.Payload["command"]
	}

	chat, err := ms.db.GetChat(payload.Message.ChatID)
	payload.Chat = chat

	if errors.Is(err, pgx.ErrNoRows) {
		payload.Chat, err = ms.db.CreateChat(payload.Message.ChatID)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	member, err := ms.db.GetChatMember(payload.Message.ChatID, payload.Message.UserID)
	payload.ChatMember = member

	if errors.Is(err, pgx.ErrNoRows) {
		user := bot.User(payload.Message.UserID)

		if user.Error.Code != 0 {
			panic(user.Error)
		}

		payload.ChatMember, err = ms.db.CreateChatMember(payload.Message.ChatID, payload.Message.UserID, user.Name)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	if object.Message.ChatID > consts.ChatsStartWith {
		if payload.Message.Text == "[club196374631|@darlingcontrol]" {
			new(commands.ControlCmd).Execute(bot, payload)
			return
		}

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
