package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

const USER_KEY = "user"

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := m.sessionStore.GetUserID(r.Context())
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if userID == uuid.Nil {
			next.ServeHTTP(w, r)
			return
		}

		user, _ := m.userRepository.GetByID(r.Context(), userID)
		if user == nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), USER_KEY, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) GetUserFromCtx(ctx context.Context) *models.User {
	user, ok := ctx.Value(USER_KEY).(*models.User)
	if !ok {
		return nil
	}

	return user
}
