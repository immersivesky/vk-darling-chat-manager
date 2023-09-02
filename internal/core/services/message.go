package services

import (
	"fmt"
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/keyboard"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
	"time"
)

type MessageService struct{}

var welcomeKeyboard = keyboard.Create().
	Text("📚 Помощь", `{"command": "help"}`).
	Text("📒 Профиль", `{"command": "profile"}`).
	Text("🏡 Чаты", `{"command": "chats"}`).
	JSON()

func NewMessageService(commandsSvc *commands.CommandsService) scene.Message {
	return func(bot *API.Bot, object update.Object) {
		object.Message.Date = int(time.Now().UnixMilli())
		fmt.Printf("New Message event = %+v\n", object.Message)
		payload := commands.NewCommandPayload(object.Message)

		if object.Message.ChatID > 2e9 {
			commandsSvc.PublicExecute(bot, payload)
		} else {
			commandsSvc.PrivateExecute(bot, payload)
		}
	}
}
