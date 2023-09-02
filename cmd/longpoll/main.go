package main

import (
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
	"gitlab.com/immersivesky/affinitycm-vk/internal/pkg/longpoll"
	"log"
	"os"
)

func main() {
	commandsService := commands.NewCommandsService()
	messageService := services.NewMessageService(commandsService)

	client := longpoll.NewLongpoll(
		os.Getenv("TOKEN"),
		messageService,
	)

	log.Fatalln(client.Run())
}
