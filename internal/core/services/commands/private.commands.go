package commands

import (
	"regexp"

	"github.com/botscommunity/vkgo/API"
)

func (c *Service) PrivateExecute(bot *API.Bot, payload *Payload) {
	for pattern, command := range c.Private {
		if locate, err := regexp.MatchString(pattern, payload.message.Text); err != nil {
			panic(err)
		} else if locate {
			command(bot, payload)
		}
	}
}
