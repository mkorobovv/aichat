package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

func (ctr *Controller) SendMessage(w http.ResponseWriter, r *http.Request) {
	var dtoIn RequestSendMessage

	err := json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		err = responseError{
			Kind:   "validation",
			Detail: err.Error(),
			status: http.StatusBadRequest,
		}

		ctr.handleError(w, err)

		return
	}

	req := dtoIn.ToRequest()

	message, err := ctr.chatUC.SendMessage(r.Context(), req)
	if err != nil {
		err = responseError{
			Kind:   "business",
			Detail: err.Error(),
			status: http.StatusInternalServerError,
		}

		ctr.handleError(w, err)

		return
	}

	dtoOut := ToResponse(message)

	err = json.NewEncoder(w).Encode(dtoOut)
	if err != nil {
		err = responseError{
			Kind:   "business",
			Detail: err.Error(),
			status: http.StatusInternalServerError,
		}

		ctr.handleError(w, err)

		return
	}
}

type RequestSendMessage struct {
	ChatID  *int64 `json:"chat_id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

func (dto *RequestSendMessage) ToRequest() chat.RequestCreateMessage {
	return chat.RequestCreateMessage{
		ChatID:  dto.ChatID,
		UserID:  dto.UserID,
		Content: dto.Content,
	}
}

type ResponseSendMessage struct {
	MessageID int64  `json:"message_id"`
	ChatID    int64  `json:"chat_id"`
	Role      string `json:"role"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func ToResponse(message chat.Message) ResponseSendMessage {
	return ResponseSendMessage{
		MessageID: message.MessageID,
		ChatID:    message.ChatID,
		Role:      message.Role,
		UserID:    message.UserID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt.Format(time.RFC3339),
	}
}
