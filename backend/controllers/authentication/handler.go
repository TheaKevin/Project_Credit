package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Login(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error bad request"})
		return
	}

	status, user, err := h.Service.Login(req)
	if err != nil {
		c.JSON(status, gin.H{"message": "Error bad request"})
	}

	c.JSON(status, gin.H{
		"message": "success",
		"user":    user,
	})
}
