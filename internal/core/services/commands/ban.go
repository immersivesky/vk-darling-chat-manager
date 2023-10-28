package commands

import (
	"fmt"
	"net/url"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"

	"gitlab.com/immersivesky/affinitycm-vk/pkg/consts"

	"github.com/botscommunity/vkgo/API"
	"github.com/botscommunity/vkgo/types"
)

type BanCmd struct{}

func (c *BanCmd) Execute(bot *API.Bot, payload *domain.Payload) {
	sent := struct {
		Response bool        `json:"response"`
		Error    types.Error `json:"error"`
	}{}
	params := url.Values{}

	params.Add("chat_id", fmt.Sprint(payload.Message.ChatID-consts.ChatsStartIn))
	params.Add("member_id", fmt.Sprint(payload.Message.Reply.UserID))

	err := bot.CallString("messages.removeChatUser", params, &sent)
	fmt.Println(sent, err)
}
