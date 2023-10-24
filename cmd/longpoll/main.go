package main

import (
	"log"
	"os"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/scenes"
	"gitlab.com/immersivesky/affinitycm-vk/internal/pkg/longpoll"
)

func main() {
	commandsService := commands.NewCommandsService()
	messageSceneService := scenes.NewMessageSceneService(commandsService)

	client := longpoll.NewLongpoll(
		os.Getenv("TOKEN"),
		messageSceneService,
	)

	log.Fatalln(client.Run())
}
