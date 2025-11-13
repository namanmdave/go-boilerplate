package service

import (
	"context"
	"database/sql"
	"go-boilerplate/dto"
	"go-boilerplate/server/ws"
	"go-boilerplate/store"

	db "go-boilerplate/store/sqlc"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type GroupService struct {
	*BaseService
}

func NewGroupService(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *GroupService {
	return &GroupService{
		BaseService: NewBaseService(store, logger, validator),
	}
}

func (gs *GroupService) CreateGroup(ctx context.Context, groupID, name, description string) (db.Group, error) {
	dbGroup, err := gs.Store.Querier.CreateGroup(ctx, db.CreateGroupParams{
		GroupID: groupID,
		Name:    name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})
	if err != nil {
		gs.Logger.Errorf("failed to create group: %v", err)
		return db.Group{}, err
	}
	return dbGroup, nil
}

func (gs *GroupService) HandleGroupMessage(ctx context.Context, group_id string, message chan dto.Message) error {
	for msg := range message {
		userID, err := gs.Store.Querier.GetUserIDInGroup(ctx, db.GetUserIDInGroupParams{
			UserID:  msg.ChatID,
			GroupID: group_id,
		})
		if err != nil && err != sql.ErrNoRows {
			gs.Logger.Errorf("failed to verify user in group: %v", err)
			continue
		}

		if userID == "" {
			gs.Store.Querier.AddUserToGroup(ctx, db.AddUserToGroupParams{
				UserID:  msg.ChatID,
				GroupID: group_id,
			})
		}

		// can have some kind of queue here so, there wont be drop in messages
		gs.Store.Querier.CreateGroupMessage(ctx, db.CreateGroupMessageParams{
			GroupID: group_id,
			SenderID: sql.NullString{
				String: msg.ChatID,
				Valid:  msg.ChatID != "",
			},
			Type:    "user",
			Message: msg.Message,
		})

		for _, conn := range ws.GroupClients[group_id].Conn {
			err := conn.WriteJSON(msg)
			if err != nil {
				gs.Logger.Errorf("failed to send message to group client: %v", err)
			}
		}
	}

	return nil
}
