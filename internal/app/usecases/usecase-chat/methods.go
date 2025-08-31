package usecase_chat

import (
	"context"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

func (uc *UseCase) CreateChat(ctx context.Context, userID int64) (int64, error) {
	chat := chat.NewChat(userID)

	id, err := uc.chatRepository.CreateChat(ctx, chat)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UseCase) SendMessage(ctx context.Context, request chat.RequestCreateMessage) (chat.Message, error) {
	if request.ChatID == nil {
		chatId, err := uc.CreateChat(ctx, request.UserID)
		if err != nil {
			return chat.Message{}, err
		}

		request.ChatID = &chatId
	}

	message := chat.NewMessage(request)

	messages, err := uc.chatRepository.GetMessages(ctx, message.ChatID)
	if err != nil {
		return chat.Message{}, err
	}

	_, err = uc.chatRepository.SaveMessage(ctx, message)
	if err != nil {
		return chat.Message{}, err
	}

	messages = append(messages, message)

	responseMessage, err := uc.aiAssistantGateway.SendMessage(ctx, messages)
	if err != nil {
		return chat.Message{}, err
	}

	id, err := uc.chatRepository.SaveMessage(ctx, responseMessage)
	if err != nil {
		return chat.Message{}, err
	}

	responseMessage.MessageID = id

	return responseMessage, nil
}
