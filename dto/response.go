package dto

import "github.com/gin-gonic/gin"

type JSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(code int, message string, data any) *JSONResponse {
	return &JSONResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(code int, message, err string) *JSONResponse {
	return &JSONResponse{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

func (r *JSONResponse) Write(c *gin.Context) {
	c.JSON(r.Code, r)
}
