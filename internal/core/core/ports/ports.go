package ports

import (
	"regexp"

	"github.com/botscommunity/vkgo/API"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/core/domain"
)

type DB interface {
	GetChat(int) (*domain.Chat, error)
	CreateChat(int) (*domain.Chat, error)
}

type Commands map[*regexp.Regexp]Command

type Command interface {
	Execute(*API.Bot, *domain.Payload)
}
