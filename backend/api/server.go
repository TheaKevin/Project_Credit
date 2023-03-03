package api

import (
	"os"
	"project_credit_sinarmas/backend/controllers/skalaAngsuran"
	"project_credit_sinarmas/backend/controllers/stagingCustomer"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}
	sa := skalaAngsuran.NewRepository(s.DB)
	sc := stagingCustomer.NewRepository(s.DB)
	c := cron.New()
	c.AddFunc("@every 30m", func() { sc.GetStagingCustomer() })
	c.AddFunc("@every 15m", func() { sa.GenerateSkalaAngsuran() })
	c.Start()

	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
