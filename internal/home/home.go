// internal/home/home.go
package home

import (
	"net/http"

	"github.com/dalemusser/gowebcore/render"
	"github.com/go-chi/chi/v5"

	"github.com/dalemusser/strata_hub/internal/handler"
)

// Home wraps the shared handler so we can access
// Config, Logger, DB client, etc., inside methods.
type Home struct {
	*handler.Handler
}

// MountRoutes registers all routes owned by the Home feature.
// Call this from internal/routes/routes.go.
func MountRoutes(r chi.Router, h *handler.Handler) {
	ctrl := &Home{Handler: h}
	r.Get("/", ctrl.Index)
}

// Index renders the landing page at "/".
func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	render.Page(w, r, "home.html", map[string]any{
		"Title": "Strata Hub",
	})
}
