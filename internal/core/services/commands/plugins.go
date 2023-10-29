package commands

import (
	"fmt"
	"github.com/botscommunity/vkgo/keyboard"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/ports"

	"github.com/botscommunity/vkgo/API"
)

type PluginsCmd struct {
	db ports.DB
}

func (c *PluginsCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	var (
		kb           = keyboard.Create().Inline()
		plugins, err = c.db.GetPlugins(1)
	)

	if err != nil {
		panic(err)
	}

	if len(plugins) > 0 {
		for i, plugin := range plugins {
			if i%2 == 0 {
				kb.Row()
			}

			var status string
			if plugin.Status == 0 {
				status = "⛔"
			}

			kb.Text(status+plugin.Name, `{"command": "badges!"}`)
		}
	} else {
		kb.Text("🗓 Список пуст")
	}

	bot.SendMessage(payload.Message.ChatID, fmt.Sprintf("%s, plugins:", payload.ChatMember.Name), kb)
}
