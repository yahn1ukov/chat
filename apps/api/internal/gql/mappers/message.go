package mappers

import (
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

func (m *Mapper) DbMessageToMessage(message *models.Message) *gqlmodels.Message {
	return &gqlmodels.Message{
		ID:        message.ID,
		UserID:    message.UserID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}
