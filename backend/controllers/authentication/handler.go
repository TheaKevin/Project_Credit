package authentication

import (
	"fmt"

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
	req.Email = c.Query("email")
	req.Password = c.Query("password")

	status, user, err := h.Service.Login(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(status, gin.H{
			"message": "Email atau password salah!",
		})
	} else {
		c.JSON(status, gin.H{
			"message": "success",
			"user":    user,
		})
	}
}
