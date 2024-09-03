package mappers

import (
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

func (m *Mapper) DbUserToUser(user *models.User) *gqlmodels.User {
	return &gqlmodels.User{
		ID:        user.ID,
		Username:  user.Username,
		Color:     user.Color,
		CreatedAt: user.CreatedAt,
	}
}

func (m *Mapper) DbUserToMessageUser(user *models.User) *gqlmodels.MessageUser {
	return &gqlmodels.MessageUser{
		ID:       user.ID,
		Username: user.Username,
		Color:    user.Color,
	}
}
