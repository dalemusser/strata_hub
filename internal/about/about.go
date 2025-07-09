package about

import (
	"net/http"

	"github.com/dalemusser/gowebcore/render"
	"github.com/dalemusser/strata_hub/internal/handler"
	"github.com/go-chi/chi/v5"
)

type About struct{ *handler.Handler }

func MountRoutes(r chi.Router, h *handler.Handler) {
	ctrl := &About{h}
	r.Get("/about", ctrl.Show)
}

func (a *About) Show(w http.ResponseWriter, r *http.Request) {
	render.Page(w, r, "about.html", map[string]any{
		"Title": "About Strata Hub",
	})
}
