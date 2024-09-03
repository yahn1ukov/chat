package gql

import (
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/directives"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/graph"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/resolvers"
	"go.uber.org/fx"
)

type Gql struct {
	*handler.Server
}

type Params struct {
	fx.In

	Resolver  *resolvers.Resolver
	Directive *directives.Directive
}

func New(p Params) *Gql {
	config := graph.Config{
		Resolvers: p.Resolver,
		Directives: graph.DirectiveRoot{
			IsAuthenticated: p.Directive.IsAuthenticated,
		},
	}

	schema := graph.NewExecutableSchema(config)

	server := handler.New(schema)

	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	server.Use(extension.Introspection{})

	return &Gql{
		server,
	}
}
