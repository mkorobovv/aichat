package config

import (
	"errors"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

func New() (config Config) {
	err := cleanenv.ReadConfig("./config.yml", &config)
	if err != nil {
		err = errors.New(strings.ReplaceAll(err.Error(), ", ", ",\n"))
		panic(err)
	}

	return config
}
