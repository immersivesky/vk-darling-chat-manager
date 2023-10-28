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
	chat := new(domain.Chat)

	if err := db.client.QueryRow(context.Background(), "SELECT chat_id FROM chats WHERE chat_id = $1;", chatID).Scan(&chat.ChatID); err != nil {
		return nil, err
	}

	return chat, nil
}

func (db *DB) CreateChat(chatID int) (*domain.Chat, error) {
	chat := new(domain.Chat)

	if err := db.client.QueryRow(context.Background(), "INSERT INTO chats(chat_id) VALUES($1) RETURNING chat_id;", chatID).Scan(&chat.ChatID); err != nil {
		return nil, err
	}

	return chat, nil
}

func (db *DB) GetChatMember(chatID, memberID int) (*domain.ChatMember, error) {
	member := new(domain.ChatMember)

	if err := db.client.QueryRow(context.Background(), "SELECT name FROM chats_members WHERE chat_id = $1 AND member_id = $2", chatID, memberID).Scan(&member.Name); err != nil {
		return nil, err
	}

	return member, nil
}

func (db *DB) CreateChatMember(chatID, memberID int, name string) (*domain.ChatMember, error) {
	member := new(domain.ChatMember)

	if err := db.client.QueryRow(
		context.Background(),
		"INSERT INTO chats_members(chat_id, member_id, name) VALUES($1, $2, $3) RETURNING name;", chatID, memberID, name).
		Scan(&member.Name); err != nil {
		return nil, err
	}

	return member, nil
}
