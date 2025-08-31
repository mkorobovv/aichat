package postgres

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const pgxDriverName = "pgx"

type Config struct {
	Host     string `env:"HOST" env-required:"true" yaml:"host"`
	Port     string `env:"PORT" env-required:"true" yaml:"port"`
	User     string `env:"USER" env-required:"true" yaml:"user"`
	Password string `env:"PASSWORD" env-required:"true" yaml:"password"`
	Name     string `env:"NAME" env-required:"true" yaml:"name"`
	TimeZone string `yaml:"timezone"`
}

func New(config Config) (*sqlx.DB, error) {
	connString := postgresConnectionString(config.Host, config.Port, config.User, config.Password, config.Name, config.TimeZone)

	db, err := sqlx.Open(pgxDriverName, connString)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func postgresConnectionString(host, port, user, password, name, timeZone string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, name, timeZone)
}
