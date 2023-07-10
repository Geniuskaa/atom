package database

import (
	"atom/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, conf *config.Entity) (*pgxpool.Pool, error) {

	//conf := pgxpool.Config{
	//	ConnConfig:            connConf,
	//	BeforeConnect:         nil,
	//	AfterConnect:          nil,
	//	BeforeAcquire:         nil,
	//	AfterRelease:          nil,
	//	BeforeClose:           nil,
	//	MaxConnLifetime:       0,
	//	MaxConnLifetimeJitter: 0,
	//	MaxConnIdleTime:       0,
	//	MaxConns:              0,
	//	MinConns:              0,
	//	HealthCheckPeriod:     0,
	//}

	poolConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.DB.User, conf.DB.Pass, conf.DB.Hostname, conf.DB.Port, conf.DB.Name))
	if err != nil {
		return nil, fmt.Errorf("NewPostgresPool failed: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("NewPostgresPool failed: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewPostgresPool failed: %w", err)
	}

	return pool, nil
}
