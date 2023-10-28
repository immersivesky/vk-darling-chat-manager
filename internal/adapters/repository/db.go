package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/immersivesky/affinitycm-vk/internal/core/domain"
)

type DB struct {
	client *pgxpool.Pool
}

func NewDB(dsn string) (*DB, error) {
	client, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &DB{
		client: client,
	}, nil
}

func (db *DB) GetChat(chatID int) (*domain.Chat, error) {
	chat := &domain.Chat{}

	if err := db.client.QueryRow(context.Background(), "SELECT chat_id FROM chats WHERE chat_id = $1;", chatID).Scan(&chat.ID); err != nil {
		return nil, err
	}

	return chat, nil
}

func (db *DB) CreateChat(chatID int) (*domain.Chat, error) {
	chat := &domain.Chat{}

	if err := db.client.QueryRow(context.Background(), "INSERT INTO chats(chat_id) VALUES($1) RETURNING chat_id;", chatID).Scan(&chat.ID); err != nil {
		return nil, err
	}

	return chat, nil
}
