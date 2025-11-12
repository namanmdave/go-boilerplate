package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler returns the health status of the server
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}

// OptionsHandler handles OPTIONS requests for CORS preflight
func OptionsHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusNoContent)
}
