package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

type Repository interface {
	Create(context.Context, uuid.UUID, *models.Message) (*models.Message, error)
	GetAll(context.Context) ([]*models.Message, error)
}
