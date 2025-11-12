package handler

import (
	"go-boilerplate/dto"
	"go-boilerplate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	*BaseHandler
	UserService *service.UserService
}

// NewUserHandler creates a new user handler with base handler and user service
func NewUserHandler(baseHandler *BaseHandler, userService *service.UserService) *UserHandler {
	return &UserHandler{
		BaseHandler: baseHandler,
		UserService: userService,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	createUserReq := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		uh.Logger.Errorw("failed to bind request", "error", err)
		dto.NewErrorResponse(http.StatusBadRequest, "invalid request", err.Error()).Write(c)
		return
	}

	if err := uh.Validator.Struct(createUserReq); err != nil {
		uh.Logger.Errorw("validation failed", "error", err)
		dto.NewErrorResponse(http.StatusBadRequest, "validation failed", err.Error()).Write(c)
		return
	}

	user, err := uh.UserService.CreateUser(c.Request.Context(), &createUserReq)
	if err != nil {
		uh.Logger.Errorw("failed to create user", "error", err)
		dto.NewErrorResponse(http.StatusInternalServerError, "failed to create user", err.Error()).Write(c)
		return
	}

	dto.NewSuccessResponse(http.StatusCreated, "user created successfully", map[string]any{
		"id":    user.ID,
		"email": user.Email,
	}).Write(c)
}
