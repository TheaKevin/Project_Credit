package transaction

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
func (h *Handler) GetTransaction(c *gin.Context) {
	transaction, status, err := h.Service.GetTransaction()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    transaction,
	})
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var req []DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error bad request"})
		return
	}

	status, err := h.Service.UpdateTransaction(req)
	if err != nil {
		c.JSON(status, gin.H{"message": "Error bad request"})
	}

	c.JSON(status, gin.H{
		"message": "success",
	})
}
