package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yahn1ukov/chat/apps/api/internal/config"
	"github.com/yahn1ukov/chat/apps/api/internal/gql"
	"github.com/yahn1ukov/chat/apps/api/internal/http/middlewares"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc         fx.Lifecycle
	Config     *config.Config
	Middleware *middlewares.Middleware
	GQLServer  *gql.Gql
}

func Run(p Params) {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", p.Middleware.Session(p.Middleware.Auth(p.Middleware.Loaders(p.GQLServer))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", p.Config.HTTP.Port),
		Handler: mux,
	}

	p.Lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
