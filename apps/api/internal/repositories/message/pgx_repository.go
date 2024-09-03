package message

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/database"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
	"go.uber.org/fx"
)

type MessagePgxRepository struct {
	database *database.Database
}

var _ Repository = (*MessagePgxRepository)(nil)

type Params struct {
	fx.In

	Database *database.Database
}

func New(p Params) *MessagePgxRepository {
	return &MessagePgxRepository{
		database: p.Database,
	}
}

var selectedFields = []string{
	"id",
	"user_id",
	"text",
	"created_at",
}

func (r *MessagePgxRepository) Create(ctx context.Context, userID uuid.UUID, message *models.Message) (*models.Message, error) {
	query, args, err := sq.
		Insert("messages").
		Columns("user_id", "text").
		Values(userID, message.Text).
		Suffix("RETURNING id, user_id, text, created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var createdMessage models.Message
	if err = r.database.
		QueryRow(ctx, query, args...).
		Scan(
			&createdMessage.ID,
			&createdMessage.UserID,
			&createdMessage.Text,
			&createdMessage.CreatedAt,
		); err != nil {
		return nil, err
	}

	return &createdMessage, nil
}

func (r *MessagePgxRepository) GetAll(ctx context.Context) ([]*models.Message, error) {
	query, args, err := sq.
		Select(selectedFields...).
		From("messages").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.database.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err = rows.Scan(
			&message.ID,
			&message.UserID,
			&message.Text,
			&message.CreatedAt,
		); err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
