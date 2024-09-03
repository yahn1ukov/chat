package user

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/dto"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/mappers"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
	userRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/user"
	"go.uber.org/fx"
)

type Service interface {
	Create(context.Context, *dto.CreateUserDto) (*gqlmodels.User, error)
	GetAll(context.Context) ([]*gqlmodels.User, error)
	GetByID(context.Context, uuid.UUID) (*gqlmodels.User, error)
}

type UserService struct {
	mapper         *mappers.Mapper
	userRepository userRepository.Repository
}

var _ Service = (*UserService)(nil)

type Params struct {
	fx.In

	Mapper         *mappers.Mapper
	UserRepository userRepository.Repository
}

func New(p Params) *UserService {
	return &UserService{
		mapper:         p.Mapper,
		userRepository: p.UserRepository,
	}
}

func (s *UserService) Create(ctx context.Context, dto *dto.CreateUserDto) (*gqlmodels.User, error) {
	if dto.Username == "" {
		return nil, ErrUsernameRequired
	}

	user := &models.User{
		Username: dto.Username,
		Color:    fmt.Sprintf("#%06x", rand.IntN(0xFFFFFF+1)),
	}

	createdUser, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	output := s.mapper.DbUserToUser(createdUser)

	return output, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]*gqlmodels.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*gqlmodels.User
	for _, user := range users {
		mappedUser := s.mapper.DbUserToUser(user)
		output = append(output, mappedUser)
	}

	return output, nil
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*gqlmodels.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	output := s.mapper.DbUserToUser(user)

	return output, nil
}
