// Package pagesections is the Page Sections togo plugin: Per-page controllable sections: admins drag-reorder and show/hide a page's blocks and save the arrangement as the public default.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package pagesections

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("page-sections", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/page-sections/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"page-sections","status":"ok"}`))
		})
		return nil
	})
}
