package user

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yahn1ukov/chat/apps/api/internal/database"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
	"go.uber.org/fx"
)

type UserPgxRepository struct {
	database *database.Database
}

var _ Repository = (*UserPgxRepository)(nil)

type Params struct {
	fx.In

	Database *database.Database
}

func New(p Params) *UserPgxRepository {
	return &UserPgxRepository{
		database: p.Database,
	}
}

var selectedFields = []string{
	"id",
	"username",
	"color",
	"created_at",
}

func (r *UserPgxRepository) FindManyByIDs(ctx context.Context, ids []uuid.UUID) ([]*models.User, error) {
	query, args, err := sq.
		Select(selectedFields...).
		From("users").
		Where(sq.Eq{"id": ids}).
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

	users := make([]*models.User, 0, len(ids))
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Color,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserPgxRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query, args, err := sq.
		Insert("users").
		Columns("username", "color").
		Values(user.Username, user.Color).
		Suffix("RETURNING id, username, color, created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var createdUser models.User
	if err = r.database.
		QueryRow(ctx, query, args...).
		Scan(
			&createdUser.ID,
			&createdUser.Username,
			&createdUser.Color,
			&createdUser.CreatedAt,
		); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrAlreadyExists
		}

		return nil, err
	}

	return &createdUser, nil
}

func (r *UserPgxRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	query, args, err := sq.
		Select(selectedFields...).
		From("users").
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

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Color,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPgxRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query, args, err := sq.
		Select(selectedFields...).
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	if err = r.database.
		QueryRow(ctx, query, args...).
		Scan(
			&user.ID,
			&user.Username,
			&user.Color,
			&user.CreatedAt,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}
