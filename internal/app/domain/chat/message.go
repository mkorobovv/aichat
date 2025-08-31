package chat

import "time"

const (
	RoleUser      = "user"
	RoleAssistant = "assistant"
)

type Chat struct {
	ChatID    int64
	UserID    int64
	CreatedAt time.Time
}

func NewChat(userID int64) Chat {
	return Chat{
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}

type Message struct {
	MessageID int64
	ChatID    int64
	UserID    int64
	Role      string
	Content   string
	CreatedAt time.Time
}

func NewMessage(req RequestCreateMessage) Message {
	return Message{
		ChatID:    *req.ChatID,
		UserID:    req.UserID,
		Role:      RoleUser,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
}

type RequestCreateMessage struct {
	ChatID  *int64 `json:"chat_id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}
