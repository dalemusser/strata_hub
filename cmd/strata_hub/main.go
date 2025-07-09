package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/dalemusser/strata_hub/internal/config"
	"github.com/dalemusser/strata_hub/internal/handler"
	"github.com/dalemusser/strata_hub/internal/routes"

	gwconfig "github.com/dalemusser/gowebcore/config"
	"github.com/dalemusser/gowebcore/server"
	"github.com/go-chi/chi/v5"
)

var cfg config.Config

func main() {
	root := &cobra.Command{
		Use:   "strata_hub",
		Short: "Central hub for the Strata platform",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return gwconfig.Load(&cfg,
				gwconfig.WithEnvPrefix("HUB"),
				gwconfig.WithConfigFile("config.toml"), // ← we fixed the extension earlier
				gwconfig.WithFlagSet(cmd.Flags()))
		},
	}

	root.AddCommand(serveCmd())

	if err := root.Execute(); err != nil {
		// Use the default logger here because the cobra root may run before we add
		// a logger to the context.
		slog.Default().Error("fatal", "err", err)
		os.Exit(1)
	}
}

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			// -----------------------------------------------------------------
			// 1) Create a base context with a logger stored in it
			// -----------------------------------------------------------------
			ctx := slog.NewContext(cmd.Context(), slog.Default())

			// -----------------------------------------------------------------
			// 2) Build feature handler (shares cfg and logger)
			// -----------------------------------------------------------------
			h := handler.New(ctx, &cfg, nil) // DB client is nil for now

			// -----------------------------------------------------------------
			// 3) Register all routes
			// -----------------------------------------------------------------
			r := chi.NewRouter()
			routes.RegisterAllRoutes(r, h)

			// -----------------------------------------------------------------
			// 4) Wrap the router with gowebcore’s http.Server
			// -----------------------------------------------------------------
			srv := server.New(cfg.Base, r)

			// -----------------------------------------------------------------
			// 5) Start serving (gowebcore will read the logger from ctx)
			// -----------------------------------------------------------------
			return server.Serve(ctx, srv, cfg.CertFile, cfg.KeyFile)
		},
	}
}
