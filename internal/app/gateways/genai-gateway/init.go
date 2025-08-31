package genai_gateway

import (
	"context"
	"log/slog"

	"github.com/mkorobovv/aichat/internal/app/config"
	"google.golang.org/genai"
)

type GenAIGateway struct {
	logger *slog.Logger
	config config.GenAIGateway
	client *genai.Client
}

func New(ctx context.Context, logger *slog.Logger, cfg config.GenAIGateway) *GenAIGateway {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.ClientSecret,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		logger.Error(err.Error(), "source", "NewGenAIGateway")

		panic(err)
	}

	return &GenAIGateway{
		logger: logger,
		config: cfg,
		client: client,
	}
}
