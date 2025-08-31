package controller

import (
	"log/slog"

	usecase_chat "github.com/mkorobovv/aichat/internal/app/usecases/usecase-chat"
)

type Controller struct {
	logger *slog.Logger
	chatUC *usecase_chat.UseCase
}

func New(logger *slog.Logger, chatUC *usecase_chat.UseCase) *Controller {
	return &Controller{
		logger: logger,
		chatUC: chatUC,
	}
}
