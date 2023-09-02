package commands

import (
	"github.com/botscommunity/vkgo/API"
	"regexp"
)

var privateCommands = map[string]command{
	pingPattern:    pingScript,
	helpPattern:    helpScript,
	profilePattern: profileScript,
}

func (c *CommandsService) PrivateExecute(bot *API.Bot, payload *Payload) {
	for pattern, command := range c.Private {
		if locate, err := regexp.MatchString(pattern, payload.message.Text); err != nil {
			panic(err)
		} else if locate {
			command(bot, payload)
		}
	}
}
