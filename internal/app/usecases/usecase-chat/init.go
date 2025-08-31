package usecase_chat

import (
	"context"
	"log/slog"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

type UseCase struct {
	logger             *slog.Logger
	aiAssistantGateway aiAssistantGateway
	chatRepository     chatRepository
}

type aiAssistantGateway interface {
	SendMessage(ctx context.Context, messages []chat.Message) (chat.Message, error)
}

type chatRepository interface {
	SaveMessage(ctx context.Context, message chat.Message) (int64, error)
	CreateChat(ctx context.Context, chat chat.Chat) (int64, error)
	GetMessages(ctx context.Context, chatId int64) ([]chat.Message, error)
}

func New(logger *slog.Logger, aiAssistantGateway aiAssistantGateway, chatRepository chatRepository) *UseCase {
	return &UseCase{
		logger:             logger,
		aiAssistantGateway: aiAssistantGateway,
		chatRepository:     chatRepository,
	}
}
