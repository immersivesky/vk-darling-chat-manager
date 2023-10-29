package ports

import (
	"github.com/botscommunity/vkgo/API"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
	"regexp"
)

type DB interface {
	GetChat(int) (*domain.Chat, error)
	CreateChat(int) (*domain.Chat, error)

	GetChatMember(int, int) (*domain.ChatMember, error)
	CreateChatMember(int, int, string) (*domain.ChatMember, error)

	GetPlugins(int) ([]*domain.Plugin, error)
}

type Commands map[*regexp.Regexp]Command

type Command interface {
	Execute(*API.Bot, *domain.Payload)
}
