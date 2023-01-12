package transaction

import (
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

type Query struct {
	branch  string
	company string
	start   time.Time
	end     time.Time
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

	company, status, err := h.Service.GetCompany()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	branch, status, err := h.Service.GetBranch()
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
		"company": company,
		"branch":  branch,
	})
}

func (h *Handler) GetTransactionFilter(c *gin.Context) {
	var q Query
	q.branch = c.Query("branch")
	q.company = c.Query("company")

	start, err := time.Parse("2006-01-02T15:04:05Z", c.Query("start"))
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	end, err := time.Parse("2006-01-02T15:04:05Z", c.Query("end"))
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	q.start = start
	q.end = end

	transaction, status, err := h.Service.GetTransactionFilter(q)
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
