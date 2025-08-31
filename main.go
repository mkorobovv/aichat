package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/mkorobovv/aichat/internal/app/config"
	genai_gateway "github.com/mkorobovv/aichat/internal/app/gateways/genai-gateway"
	openai_gateway "github.com/mkorobovv/aichat/internal/app/gateways/openai-gateway"
	http_adapter "github.com/mkorobovv/aichat/internal/app/http-adapter"
	chat_repository "github.com/mkorobovv/aichat/internal/app/repositories/chat-repository"
	usecase_chat "github.com/mkorobovv/aichat/internal/app/usecases/usecase-chat"
	"github.com/mkorobovv/aichat/internal/pkg/logger"
	"github.com/mkorobovv/aichat/internal/pkg/postgres"
	"golang.org/x/sync/errgroup"
)

func main() {
	l := logger.New()

	cfg := config.New()

	chatDB, err := postgres.New(cfg.Databases.Chat)
	if err != nil {
		slog.Error(err.Error(), "source", "postgres")

		panic(err)
	}
	defer func() {
		err := chatDB.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chatRepository := chat_repository.New(l, chatDB)

	_ = openai_gateway.New(l, cfg.Gateways.OpenAIGateway)
	genAIGateway := genai_gateway.New(ctx, l, cfg.Gateways.GenAIGateway)

	chatUC := usecase_chat.New(l, genAIGateway, chatRepository)

	httpAdapter := http_adapter.New(l, cfg.HttpAdapter, chatUC)

	err = start(
		ctx,
		httpAdapter,
	)
	if err != nil {
		slog.Error(err.Error(), "source", "start")
	}
}

func start(ctx context.Context, starters ...starter) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, s := range starters {
		f := func() error {
			err := s.Start(ctx)
			if err != nil {
				log.Println(err)

				log.Println("starting graceful shutdown")

				return err
			}

			return nil
		}

		g.Go(f)
	}

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}

type starter interface {
	Start(ctx context.Context) error
}
