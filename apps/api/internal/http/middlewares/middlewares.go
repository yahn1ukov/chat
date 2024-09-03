package middlewares

import (
	"github.com/yahn1ukov/chat/apps/api/internal/gql/loaders"
	sessionstore "github.com/yahn1ukov/chat/apps/api/internal/http/session_store"
	userRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/user"
	"go.uber.org/fx"
)

type Middleware struct {
	sessionStore   *sessionstore.SessionStore
	loader         *loaders.Loader
	userRepository userRepository.Repository
}

type Params struct {
	fx.In

	SessionStore   *sessionstore.SessionStore
	Loader         *loaders.Loader
	UserRepository userRepository.Repository
}

func New(p Params) *Middleware {
	return &Middleware{
		sessionStore:   p.SessionStore,
		loader:         p.Loader,
		userRepository: p.UserRepository,
	}
}
