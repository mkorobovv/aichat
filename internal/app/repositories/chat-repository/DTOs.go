package chat_repository

import (
	"time"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

type GetMessagesDTO struct {
	MessageID int64  `db:"message_id"`
	ChatID    int64  `db:"chat_id"`
	UserID    int64  `db:"user_id"`
	Role      string `db:"role"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}

func (dto *GetMessagesDTO) ToEntity() (chat.Message, error) {
	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		return chat.Message{}, err
	}

	return chat.Message{
		MessageID: dto.MessageID,
		ChatID:    dto.ChatID,
		UserID:    dto.UserID,
		Role:      dto.Role,
		Content:   dto.Content,
		CreatedAt: createdAt,
	}, nil
}
