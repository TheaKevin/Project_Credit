package report

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
	status  string
	start   time.Time
	end     time.Time
}

func (h *Handler) GetReport(c *gin.Context) {
	report, status, err := h.Service.GetReport()
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
		"data":    report,
		"company": company,
		"branch":  branch,
	})
}

func (h *Handler) GetReportFilter(c *gin.Context) {
	var q Query
	q.branch = c.Query("branch")
	q.company = c.Query("company")
	q.status = c.Query("status")

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

	report, status, err := h.Service.GetReportFilter(q)
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    report,
	})
}
