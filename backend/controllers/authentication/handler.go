package authentication

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	status, token, err := h.Service.Login(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		cookie := http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		}

		c.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Sub(time.Now().UTC().Round(time.Second)).Seconds()), "/", "", cookie.Secure, cookie.HttpOnly)

		c.JSON(status, gin.H{
			"message": "success",
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
			"message": "Password lama berbeda dengan password yang sedang digunakan!",
		})
	} else {
		c.JSON(status, gin.H{
			"message": "success",
		})
	}
}

func (h *Handler) AuthenticateUser(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	status, user, err := h.Service.AuthenticateUser(cookie)
	if err != nil {
		fmt.Println(err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, gin.H{
			"user": user,
		})
	}
}

func (h *Handler) Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	c.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Sub(time.Now().UTC().Round(time.Second)).Seconds()), "/", "", cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
