package middlewares

import (
	"context"
	"net/http"

	sessionstore "github.com/yahn1ukov/chat/apps/api/internal/http/session_store"
)

const (
	RESPONSE_WRITER_KEY = "responseWriter"
	REQUEST_KEY         = "request"
)

func (m *Middleware) Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := m.sessionStore.Get(r, sessionstore.SESSION_KEY)

		ctx := context.WithValue(r.Context(), RESPONSE_WRITER_KEY, w)
		ctx = context.WithValue(ctx, REQUEST_KEY, r)
		ctx = context.WithValue(ctx, sessionstore.SESSION_KEY, session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
