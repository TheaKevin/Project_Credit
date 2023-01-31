package authentication

import (
	"fmt"
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

func (h *Handler) ChangePassword(c *gin.Context) {
	var pass Password

	if err := c.ShouldBindJSON(&pass); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error bad request"})
		return
	}

	status, err := h.Service.ChangePassword(pass)
	if err != nil {
		fmt.Println(err)
		c.JSON(status, gin.H{
			"message": "Password baru sama dengan password lama!",
		})
	} else {
		c.JSON(status, gin.H{
			"message": "success",
		})
	}
}
