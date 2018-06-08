package go_common

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string
}

func Health(c *gin.Context) {

	c.JSON(http.StatusOK, HealthResponse{Status: "UP"})
}