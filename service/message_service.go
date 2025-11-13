package service

import (
	"context"
	"database/sql"
	"go-boilerplate/store"
	db "go-boilerplate/store/sqlc"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type MessageService struct {
	*BaseService
}

func NewMessageService(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *MessageService {
	return &MessageService{
		BaseService: NewBaseService(store, logger, validator),
	}
}

func (ms *MessageService) SendMessageToGroup(ctx context.Context, groupID, senderID, msgType, message string) (db.GroupMessage, error) {
	dbMessage, err := ms.Store.Querier.CreateGroupMessage(ctx, db.CreateGroupMessageParams{
		GroupID: groupID,
		SenderID: sql.NullString{
			String: senderID,
			Valid:  senderID != "",
		},
		Type:    msgType,
		Message: message,
	})
	if err != nil {
		ms.Logger.Errorf("failed to send message to group: %v", err)
		return db.GroupMessage{}, err
	}
	return dbMessage, nil
}

func (ms *MessageService) GetMessagesByGroupID(ctx context.Context, groupID string, limit, offset int32) ([]db.GroupMessage, error) {
	messages, err := ms.Store.Querier.GetGroupMessagesByGroupID(ctx, db.GetGroupMessagesByGroupIDParams{
		GroupID: groupID,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		ms.Logger.Errorf("failed to get messages for group %s: %v", groupID, err)
		return nil, err
	}
	return messages, nil
}
