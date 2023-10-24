package commands

import (
	"github.com/botscommunity/vkgo/API"
)

type command func(bot *API.Bot, payload *Payload)
type commands map[string]command

type Service struct {
	Private commands
	Public  commands
}

func NewCommandsService() *Service {
	return &Service{
		Private: map[string]command{
			"^(?:ping)$": pingScript,
			"^(?:help|menu|помощь|команды)$": helpScript,
			"^(profile|профиль)$":            profileScript,
		},
		Public: map[string]command{
			"^(eban)$": banScript,
		},
	}
}
