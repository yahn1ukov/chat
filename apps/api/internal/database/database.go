package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yahn1ukov/chat/apps/api/internal/config"
	"go.uber.org/fx"
)

type Database struct {
	*pgxpool.Pool
}

type Params struct {
	fx.In

	Config *config.Config
}

func New(p Params) (*Database, error) {
	pool, err := pgxpool.New(context.Background(), p.Config.DB.Postgres.URL)
	if err != nil {
		return nil, err
	}

	return &Database{
		pool,
	}, nil
}
