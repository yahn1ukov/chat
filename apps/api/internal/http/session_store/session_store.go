package sessionstore

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/yahn1ukov/chat/apps/api/internal/config"
	"go.uber.org/fx"
)

const (
	SESSION_KEY = "session"
	USER_ID_KEY = "user_id"
)

type SessionStore struct {
	*sessions.CookieStore
}

type Params struct {
	fx.In

	Config *config.Config
}

func New(p Params) *SessionStore {
	sessionAge := int((31 * 24 * time.Hour).Seconds())

	store := sessions.NewCookieStore([]byte(p.Config.Session.Secret))

	store.Options.MaxAge = sessionAge
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true

	return &SessionStore{
		store,
	}
}

func (s *SessionStore) GetSession(ctx context.Context) *sessions.Session {
	return ctx.Value(SESSION_KEY).(*sessions.Session)
}

func (s *SessionStore) GetUserID(ctx context.Context) (uuid.UUID, error) {
	session := s.GetSession(ctx)

	userID := session.Values[USER_ID_KEY]
	if userID == nil {
		return uuid.Nil, fmt.Errorf("user id not found in session")
	}

	castedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		return uuid.Nil, fmt.Errorf("user id not uuid")
	}

	return castedUserID, nil
}
