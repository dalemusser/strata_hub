// internal/handler/handler.go
package handler

import (
	"context"
	"log/slog"

	"github.com/dalemusser/strata_hub/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is injected into every feature package so all features share
// config, a logger, the DB client, and the base context.
type Handler struct {
	Ctx context.Context
	Cfg *config.Config
	Log *slog.Logger
	DB  *mongo.Client
}

// New builds the shared handler.
//
// For now we always use slog.Default(). Later—when you need per‑request
// loggers—you can add your own helper that extracts a logger from the
// context and falls back to slog.Default().
func New(ctx context.Context, cfg *config.Config, db *mongo.Client) *Handler {
	return &Handler{
		Ctx: ctx,
		Cfg: cfg,
		DB:  db,
		Log: slog.Default(),
	}
}
