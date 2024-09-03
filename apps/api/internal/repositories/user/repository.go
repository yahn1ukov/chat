package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
)

type Repository interface {
	FindManyByIDs(context.Context, []uuid.UUID) ([]*models.User, error)
	Create(context.Context, *models.User) (*models.User, error)
	GetAll(context.Context) ([]*models.User, error)
	GetByID(context.Context, uuid.UUID) (*models.User, error)
}
