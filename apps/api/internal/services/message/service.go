package message

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/yahn1ukov/chat/apps/api/internal/dto"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/gqlmodels"
	"github.com/yahn1ukov/chat/apps/api/internal/gql/mappers"
	"github.com/yahn1ukov/chat/apps/api/internal/models"
	messageRepository "github.com/yahn1ukov/chat/apps/api/internal/repositories/message"
	"go.uber.org/fx"
)

type Service interface {
	Create(context.Context, uuid.UUID, *dto.CreateMessageDto) (*gqlmodels.Message, error)
	GetAll(context.Context) ([]*gqlmodels.Message, error)
}

type MessageService struct {
	mapper            *mappers.Mapper
	messageRepository messageRepository.Repository
}

var _ Service = (*MessageService)(nil)

type Params struct {
	fx.In

	Mapper            *mappers.Mapper
	MessageRepository messageRepository.Repository
}

func New(p Params) *MessageService {
	return &MessageService{
		mapper:            p.Mapper,
		messageRepository: p.MessageRepository,
	}
}

func (s *MessageService) Create(ctx context.Context, userID uuid.UUID, dto *dto.CreateMessageDto) (*gqlmodels.Message, error) {
	if dto.Text == "" {
		return nil, ErrTextRequired
	}

	text := strings.ReplaceAll(dto.Text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")

	if utf8.RuneCountInString(text) > 200 {
		return nil, ErrTextLimit
	}

	message := &models.Message{
		Text: text,
	}

	createdMessage, err := s.messageRepository.Create(ctx, userID, message)
	if err != nil {
		return nil, err
	}

	output := s.mapper.DbMessageToMessage(createdMessage)

	return output, nil
}

func (s *MessageService) GetAll(ctx context.Context) ([]*gqlmodels.Message, error) {
	messages, err := s.messageRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*gqlmodels.Message
	for _, message := range messages {
		mappedMessage := s.mapper.DbMessageToMessage(message)
		output = append(output, mappedMessage)
	}

	return output, nil
}
