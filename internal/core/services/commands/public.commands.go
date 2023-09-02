package commands

import (
	"github.com/botscommunity/vkgo/API"
	"regexp"
)

var publicCommands = map[string]command{
	banPattern: banScript,
}

func (c *CommandsService) PublicExecute(bot *API.Bot, payload *Payload) {
	for pattern, command := range c.Public {
		if locate, err := regexp.MatchString(pattern, payload.message.Text); err != nil {
			panic(err)
		} else if locate {
			command(bot, payload)
		}
	}
}
