package app

import (
	"github.com/yahn1ukov/chat/apps/api/internal/config"
	"github.com/yahn1ukov/chat/apps/api/internal/database"
	"github.com/yahn1ukov/chat/apps/api/internal/gql"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/directives"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/loaders"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/mappers"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/resolvers"
	"github.com/yahn1ukov/chat/apps/api/internal/http"
	"github.com/yahn1ukov/chat/apps/api/internal/http/middlewares"
	sessionstore "github.com/yahn1ukov/chat/apps/api/internal/http/session_store"
	"github.com/yahn1ukov/chat/apps/api/internal/pubsub"
	messageRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/message"
	userRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/user"
	messageService "github.com/yahn1ukov/chat/apps/api/internal/services/message"
	userService "github.com/yahn1ukov/chat/apps/api/internal/services/user"
	"go.uber.org/fx"
)

func New(configPath string) *fx.App {
	return fx.New(
		fx.Provide(
			func() (*config.Config, error) {
				return config.New(configPath)
			},
			database.New,
			pubsub.New,
			sessionstore.New,
		),

		fx.Provide(
			fx.Annotate(userRepository.New, fx.As(new(userRepository.Repository))),
			fx.Annotate(messageRepository.New, fx.As(new(messageRepository.Repository))),
		),

		fx.Provide(
			fx.Annotate(userService.New, fx.As(new(userService.Service))),
			fx.Annotate(messageService.New, fx.As(new(messageService.Service))),
		),

		fx.Provide(
			middlewares.New,
			mappers.New,
			loaders.New,
			directives.New,
			resolvers.New,
			gql.New,
		),

		fx.Invoke(http.Run),
	)
}
