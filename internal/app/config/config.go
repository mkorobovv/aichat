package config

import (
	"time"

	"github.com/mkorobovv/aichat/internal/pkg/postgres"
)

type Config struct {
	Gateways    Gateways    `yaml:"gateways"`
	Databases   Databases   `yaml:"databases"`
	HttpAdapter HttpAdapter `yaml:"httpAdapter"`
}

type Gateways struct {
	OpenAIGateway OpenAIGateway `yaml:"openAIGateway"`
	GenAIGateway  GenAIGateway  `yaml:"genAIGateway"`
}

type OpenAIGateway struct {
	ClientSecret string `env:"OPENAI_CLIENT_SECRET" env-required:"true" yaml:"clientSecret"`
}

type GenAIGateway struct {
	ClientSecret string `env:"GENAI_CLIENT_SECRET" env-required:"true" yaml:"clientSecret"`
}

type Databases struct {
	Chat postgres.Config `env-prefix:"CHAT_" yaml:"chat"`
}

type HttpAdapter struct {
	Server Server `yaml:"server"`
	Router Router `yaml:"router"`
}

type Router struct {
	Shutdown Shutdown `yaml:"shutdown"`
	Timeout  Timeout  `yaml:"timeout"`
}

type Shutdown struct {
	Duration time.Duration `yaml:"duration"`
}

type Timeout struct {
	Duration time.Duration `yaml:"duration"`
}

type Server struct {
	Port              string
	Name              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
}
