package directives

import (
	"github.com/yahn1ukov/chat/apps/api/internal/http/middlewares"
	"go.uber.org/fx"
)

type Directive struct {
	middleware *middlewares.Middleware
}

type Params struct {
	fx.In

	Middleware *middlewares.Middleware
}

func New(p Params) *Directive {
	return &Directive{
		middleware: p.Middleware,
	}
}
