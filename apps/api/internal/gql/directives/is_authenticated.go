package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func (d *Directive) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := d.middleware.GetUserFromCtx(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	return next(ctx)
}
