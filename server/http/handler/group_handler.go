package handler

import (
	"go-boilerplate/dto"
	"go-boilerplate/server/ws"
	"go-boilerplate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	*BaseHandler

	GroupService *service.GroupService
}

func NewGroupHandler(baseHandler *BaseHandler, groupService *service.GroupService) *GroupHandler {
	return &GroupHandler{
		BaseHandler:  baseHandler,
		GroupService: groupService,
	}
}

func (h *GroupHandler) UpgradeHandler(c *gin.Context) {

	groupID := c.Param("group_id")
	if groupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing group_id"})
		return
	}

	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	if _, ok := ws.GroupClients[groupID]; !ok {
		h.Logger.Infof("Starting message handler for group %s", groupID)
		go h.GroupService.HandleGroupMessage(c.Request.Context(), groupID, ws.Broadcast)
	}

	// bunch of group validation
	ws.GroupClients[groupID] = dto.GroupConnection{
		Conn:    append(ws.GroupClients[groupID].Conn, conn),
		GroupID: groupID,
	}

	for {
		var msg dto.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			h.Logger.Error("error while reading the message", err)
			delete(ws.GroupClients, groupID)
			return
		}

		h.Logger.Infof("Received message in group %s: %v", groupID, msg)

		ws.Broadcast <- msg
	}
}
