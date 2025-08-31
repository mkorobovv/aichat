package openai_gateway

import (
	"context"
	"time"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

func (gw *OpenAIGateway) SendMessage(ctx context.Context, messages []chat.Message) (chat.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := buildRequest(messages)

	response, err := gw.client.Chat.Completions.New(ctx, req)
	if err != nil {
		return chat.Message{}, err
	}

	message := toEntity(messages[0].UserID, messages[0].ChatID, response)

	return message, nil
}

func buildRequest(messages []chat.Message) openai.ChatCompletionNewParams {
	req := openai.ChatCompletionNewParams{
		Model:    shared.ChatModelGPT4oMini,
		Seed:     openai.Int(1),
		Messages: make([]openai.ChatCompletionMessageParamUnion, 0, len(messages)+1),
	}

	req.Messages = append(req.Messages, openai.SystemMessage("Answer with maximum 2 sentences"))

	for _, message := range messages {
		switch message.Role {
		case chat.RoleUser:
			req.Messages = append(req.Messages, openai.UserMessage(message.Content))
		case chat.RoleAssistant:
			req.Messages = append(req.Messages, openai.AssistantMessage(message.Content))
		default:
		}
	}

	return req
}

func toEntity(userID, chatID int64, completion *openai.ChatCompletion) chat.Message {
	return chat.Message{
		UserID:    userID,
		ChatID:    chatID,
		Role:      chat.RoleAssistant,
		Content:   completion.Choices[0].Message.Content,
		CreatedAt: time.Now(),
	}
}
