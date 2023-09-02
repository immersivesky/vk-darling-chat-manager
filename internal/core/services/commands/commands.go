package commands

import (
	"github.com/botscommunity/vkgo/API"
)

type command func(bot *API.Bot, payload *Payload)
type commands map[string]command

type CommandsService struct {
	Private commands
	Public  commands
}

func NewCommandsService() *CommandsService {
	return &CommandsService{
		Private: privateCommands,
		Public:  publicCommands,
	}
}
