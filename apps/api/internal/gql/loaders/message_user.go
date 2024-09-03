package loaders

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

func GetMessageUserByID(ctx context.Context, userID uuid.UUID) (*gqlmodels.MessageUser, error) {
	loader := GetLoaderForRequest(ctx)
	return loader.messageUserLoader.Load(ctx, userID)
}

func GetMessageUsersByIDs(ctx context.Context, userIDs []uuid.UUID) ([]*gqlmodels.MessageUser, error) {
	loader := GetLoaderForRequest(ctx)
	return loader.messageUserLoader.LoadAll(ctx, userIDs)
}

func (l *Loader) getMessageUsersByIds(ctx context.Context, ids []uuid.UUID) ([]*gqlmodels.MessageUser, []error) {
	users, err := l.userRepository.FindManyByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	messageUsers := make([]*gqlmodels.MessageUser, 0, len(users))

	for _, id := range ids {
		user, ok := lo.Find(
			users,
			func(item *models.User) bool {
				return item.ID == id
			},
		)
		if !ok {
			messageUsers = append(messageUsers, nil)
		} else {
			mappedMessageUser := l.mapper.DbUserToMessageUser(user)
			messageUsers = append(messageUsers, mappedMessageUser)
		}
	}

	return messageUsers, nil
}
