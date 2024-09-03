package resolvers

import (
	sessionstore "github.com/yahn1ukov/chat/apps/api/internal/http/session_store"
	"github.com/yahn1ukov/chat/apps/api/internal/pubsub"
	messageService "github.com/yahn1ukov/chat/apps/api/internal/services/message"
	userService "github.com/yahn1ukov/chat/apps/api/internal/services/user"
	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	sessionStore   *sessionstore.SessionStore
	pubsub         *pubsub.PubSub
	userService    userService.Service
	messageService messageService.Service
}

type Params struct {
	fx.In

	SessionStore   *sessionstore.SessionStore
	PubSub         *pubsub.PubSub
	UserService    userService.Service
	MessageService messageService.Service
}

func New(p Params) *Resolver {
	return &Resolver{
		sessionStore:   p.SessionStore,
		pubsub:         p.PubSub,
		userService:    p.UserService,
		messageService: p.MessageService,
	}
}
