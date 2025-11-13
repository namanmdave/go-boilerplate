package handler

import (
	"fmt"
	"go-boilerplate/dto"
	"go-boilerplate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	*BaseHandler

	MessageService *service.MessageService
}

func NewMessageHandler(baseHandler *BaseHandler, messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{
		BaseHandler:    baseHandler,
		MessageService: messageService,
	}
}

func (mh *MessageHandler) FetchGroupMessages(c *gin.Context) {
	groupID := c.Param("group_id")

	if groupID == "" {
		mh.Logger.Errorw("group_id is required")
		dto.NewErrorResponse(http.StatusBadRequest, "group_id is required", "group_id parameter is missing").Write(c)
		return
	}

	var limitInt32 int32 = 20
	var offsetInt32 int32 = 0

	limit, _ := c.GetQuery("limit")
	if limit != "" {
		limitInt32 = parseInt32(limit)
	}

	offset, _ := c.GetQuery("offset")
	if offset != "" {
		offsetInt32 = parseInt32(offset)
	}

	messages, err := mh.MessageService.GetMessagesByGroupID(c.Request.Context(), groupID, limitInt32, offsetInt32)
	if err != nil {
		mh.Logger.Errorw("failed to fetch messages", "error", err)
		dto.NewErrorResponse(http.StatusInternalServerError, "failed to fetch messages", err.Error()).Write(c)
		return
	}

	c.JSON(200, gin.H{"messages": messages})
}

func parseInt32(s string) int32 {
	var i int32
	_, err := fmt.Sscan(s, &i)
	if err != nil {
		return 0
	}
	return i
}
