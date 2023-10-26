package main

import (
	"log"
	"os"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/callback"
	"github.com/botscommunity/vkgo/scene"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/immersivesky/affinitycm-vk/internal/adapters/repository"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/commands"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/services/scenes"
)

func main() {
	db, err := repository.NewDB("postgresql://affinitydb:huiowte4@localhost:8032/vk")
	if err != nil {
		panic(err)
	}

	publicCommandsService, privateCommandsService := commands.NewCommandsService()
	messageSceneService := scenes.NewMessageSceneService(db, publicCommandsService, privateCommandsService)

	bot := API.Create(os.Getenv("TOKEN"))

	scenes := scene.Create().
		Message(messageSceneService.GetMessageScene)

	cb := callback.Create(bot, scenes)

	app := fiber.New()
	app.Post("/callback/:confirmation", cb.Fiber)

	log.Fatalln(app.Listen(":3032"))
}
