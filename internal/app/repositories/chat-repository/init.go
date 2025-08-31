package chat_repository

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	logger *slog.Logger
	DB     *sqlx.DB
}

func New(logger *slog.Logger, db *sqlx.DB) *ChatRepository {
	return &ChatRepository{
		logger: logger,
		DB:     db,
	}
}
