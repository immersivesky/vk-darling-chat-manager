package main

import (
	"log"
	"os"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/callback"
	"github.com/botscommunity/vkgo/scene"
	fiber "github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/middleware/logger".
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

	cb := callback.Create(bot, scenes)

	app := fiber.New()

	prometheus := fiberprometheus.New("vk-affinity-callback")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// app.Use(logger.New())
	app.Post("/callback/:confirmation", cb.Fiber)

	log.Fatalln(app.Listen(":3032"))
}
