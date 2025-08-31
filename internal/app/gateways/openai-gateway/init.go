package openai_gateway

import (
	"log/slog"

	"github.com/mkorobovv/aichat/internal/app/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIGateway struct {
	logger *slog.Logger
	config config.OpenAIGateway
	client *openai.Client
}

func New(logger *slog.Logger, cfg config.OpenAIGateway) *OpenAIGateway {
	client := openai.NewClient(
		option.WithAPIKey(cfg.ClientSecret),
		option.WithMaxRetries(3),
	)

	return &OpenAIGateway{
		logger: logger,
		config: cfg,
		client: &client,
	}
}
