package genai_gateway

import (
	"context"
	"time"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
	"google.golang.org/genai"
)

func (gw *GenAIGateway) SendMessage(ctx context.Context, messages []chat.Message) (chat.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := buildRequest(messages)

	response, err := gw.client.Models.GenerateContent(ctx, "gemini-2.0-flash", req, nil)
	if err != nil {
		return chat.Message{}, err
	}

	message := toEntity(messages[0].UserID, messages[0].ChatID, response)

	return message, nil
}

func buildRequest(messages []chat.Message) []*genai.Content {
	req := []*genai.Content{
		{
			Role:  genai.RoleUser,
			Parts: []*genai.Part{{Text: "Answer with maximum 2 sentences"}},
		},
	}

	for _, m := range messages {
		switch m.Role {
		case chat.RoleUser:
			req = append(req, &genai.Content{
				Role:  genai.RoleUser,
				Parts: []*genai.Part{{Text: m.Content}},
			})
		case chat.RoleAssistant:
			req = append(req, &genai.Content{
				Role:  genai.RoleModel,
				Parts: []*genai.Part{{Text: m.Content}},
			})
		}
	}

	return req
}

func toEntity(userID, chatID int64, resp *genai.GenerateContentResponse) chat.Message {
	return chat.Message{
		UserID:    userID,
		ChatID:    chatID,
		Role:      chat.RoleAssistant,
		Content:   resp.Text(),
		CreatedAt: time.Now(),
	}
}
