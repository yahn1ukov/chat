package loaders

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vikstrous/dataloadgen"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/mappers"
	userRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/user"
	"go.uber.org/fx"
)

type ctxKey string

const KEY = ctxKey("dataloaders")

type Loader struct {
	mapper         *mappers.Mapper
	userRepository userRepository.Repository

	messageUserLoader *dataloadgen.Loader[uuid.UUID, *gqlmodels.MessageUser]
}

type Params struct {
	fx.In

	Mapper         *mappers.Mapper
	UserRepository userRepository.Repository
}

func New(p Params) *Loader {
	loader := &Loader{
		mapper:         p.Mapper,
		userRepository: p.UserRepository,
	}

	loader.messageUserLoader = dataloadgen.NewLoader(
		loader.getMessageUsersByIds,
		dataloadgen.WithWait(time.Millisecond),
	)

	return loader
}

func GetLoaderForRequest(ctx context.Context) *Loader {
	return ctx.Value(KEY).(*Loader)
}
