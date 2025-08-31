package chat_repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

func (repo *ChatRepository) CreateChat(ctx context.Context, chat chat.Chat) (id int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query, args, err := queryCreateChat(chat)
	if err != nil {
		return 0, err
	}

	err = repo.DB.GetContext(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func queryCreateChat(chat chat.Chat) (string, []interface{}, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Insert("chats.chats").
		Columns(
			"user_id",
		).
		Values(
			chat.UserID,
		).
		Suffix("RETURNING chats.chat_id")

	return query.ToSql()
}

func (repo *ChatRepository) SaveMessage(ctx context.Context, message chat.Message) (id int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query, args, err := querySaveMessage(message)
	if err != nil {
		return 0, err
	}

	err = repo.DB.GetContext(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func querySaveMessage(message chat.Message) (string, []interface{}, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Insert("chats.messages").
		Columns(
			"chat_id",
			"user_id",
			"role",
			"content",
		).
		Values(
			message.ChatID,
			message.UserID,
			message.Role,
			message.Content,
		).
		Suffix("RETURNING message_id")

	return query.ToSql()
}

func (repo *ChatRepository) GetMessages(ctx context.Context, chatID int64) ([]chat.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query, args, err := queryGetMessages(chatID)
	if err != nil {
		return nil, err
	}

	var dto []GetMessagesDTO

	err = repo.DB.SelectContext(ctx, &dto, query, args...)
	if err != nil {
		return nil, err
	}

	messages := make([]chat.Message, 0, len(dto))

	for _, d := range dto {
		message, err := d.ToEntity()
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func queryGetMessages(chatID int64) (string, []interface{}, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Select(
			"m.message_id",
			"m.chat_id",
			"m.user_id",
			"m.role",
			"m.content",
			"m.created_at",
		).
		From("chats.messages m").
		Where(sq.Eq{"m.chat_id": chatID}).
		OrderBy("m.message_id ASC").
		Limit(10)

	return query.ToSql()
}
