package skalaAngsuran

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
func (h *Handler) GenerateSkalaAngsuran(c *gin.Context) {
	sa, status, err := h.Service.GenerateSkalaAngsuran()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    sa,
	})
}
