package db

import (
	"context"

	"github.com/hinccvi/go-ddd/internal/config"
	"github.com/hinccvi/go-ddd/internal/entity"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func Connect(cfg *config.Config, log *zap.Logger) (entity.DBTX, error) {
	config, err := pgxpool.ParseConfig(cfg.Dsn)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Logger = zapadapter.NewLogger(log)

	pgx, err := pgxpool.ConnectConfig(context.TODO(), config)
	if err != nil {
		return nil, err
	}

	return pgx, nil
}
