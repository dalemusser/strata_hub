package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/dalemusser/strata_hub/internal/about"
	"github.com/dalemusser/strata_hub/internal/handler"
	"github.com/dalemusser/strata_hub/internal/home"
)

// RegisterAllRoutes is the single fan‑in invoked from main().
func RegisterAllRoutes(r chi.Router, h *handler.Handler) {
	home.MountRoutes(r, h)
	about.MountRoutes(r, h)

	// ⬆️ add future features here: auth.MountRoutes(r,h), dashboard.Mount…, etc.
}
