package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/furniture/config"
)

func ConnectToDB(cfg *config.DB) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db ping timeout: %w", err)
	}

	return db, nil
}

func connStr(cfg *config.DB) string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%d/%v?sslmode=%v", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,  cfg.SSLMode)
}