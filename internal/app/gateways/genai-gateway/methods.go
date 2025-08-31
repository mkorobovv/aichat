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

	req, cfg := buildRequest(messages)

	response, err := gw.client.Models.GenerateContent(ctx, "gemini-2.0-flash", req, cfg)
	if err != nil {
		return chat.Message{}, err
	}

	message := toEntity(messages[0].UserID, messages[0].ChatID, response)

	return message, nil
}

func buildRequest(messages []chat.Message) ([]*genai.Content, *genai.GenerateContentConfig) {
	requests := make([]*genai.Content, 0, len(messages))

	for _, m := range messages {
		switch m.Role {
		case chat.RoleUser:
			requests = append(requests, &genai.Content{
				Role:  genai.RoleUser,
				Parts: []*genai.Part{{Text: m.Content}},
			})
		case chat.RoleAssistant:
			requests = append(requests, &genai.Content{
				Role:  genai.RoleModel,
				Parts: []*genai.Part{{Text: m.Content}},
			})
		}
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: "Answer with 2 sentences maximum."},
			},
		},
	}

	return requests, config
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
