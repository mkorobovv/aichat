package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mkorobovv/aichat/internal/app/domain/chat"
)

func (ctr *Controller) SendMessage(w http.ResponseWriter, r *http.Request) {
	var dtoIn RequestSendMessage

	err := json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		ctr.handleError(w, err)

		return
	}

	req := dtoIn.ToRequest()

	message, err := ctr.chatUC.SendMessage(r.Context(), req)
	if err != nil {
		ctr.handleError(w, err)

		return
	}

	err = json.NewEncoder(w).Encode(message)
	if err != nil {
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
