package commands

import (
	"fmt"
	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/types"
	"net/url"
)

var banPattern = "^(eban|)$"

func banScript(bot *API.Bot, payload *Payload) {
	sent := struct {
		Response bool        `json:"response"`
		Error    types.Error `json:"error"`
	}{}
	params := url.Values{}

	params.Add("chat_id", fmt.Sprint(payload.message.ChatID-2e9))
	params.Add("user_id", fmt.Sprint(payload.message.Reply.UserID))

	err := bot.CallString("messages.removeChatUser", params, &sent)
	fmt.Println(sent, err)
}
