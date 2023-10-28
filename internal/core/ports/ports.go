package ports

import (
	"regexp"

	"github.com/botscommunity/vkgo/API"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
)

type DB interface {
	GetChat(int) (*domain.Chat, error)
	CreateChat(int) (*domain.Chat, error)

	GetChatMember(int, int) (*domain.ChatMember, error)
	CreateChatMember(int, int, string) (*domain.ChatMember, error)
}

type Commands map[*regexp.Regexp]Command

type Command interface {
	Execute(*API.Bot, *domain.Payload)
}
