package main

import (
	"log"
	"os"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/longpoll"
	"github.com/botscommunity/vkgo/scene"
	"gitlab.com/immersivesky/affinitycm-vk/internal/adapters/repository"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/scenes"
)

func main() {
	db, err := repository.NewDB(os.Getenv("POSTGRESQL_URI"))
	if err != nil {
		panic(err)
	}

	publicCommandsService, privateCommandsService := commands.NewCommandsService()
	messageSceneService := scenes.NewMessageSceneService(db, publicCommandsService, privateCommandsService)

	bot := API.Create(os.Getenv("TOKEN"))

	scenes := scene.Create().
		Message(messageSceneService.GetMessageScene)

	log.Fatalln(longpoll.Create(bot, scenes).Listen())
}
