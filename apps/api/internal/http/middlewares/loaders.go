package middlewares

import (
	"context"
	"net/http"

	"github.com/yahn1ukov/chat/apps/api/internal/gql/loaders"
)

func (m *Middleware) Loaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loaders.KEY, m.loader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
