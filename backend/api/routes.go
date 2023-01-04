package api

import (
	"project_credit_sinarmas/backend/controllers/stagingCustomer"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
	}))

	scRepo := stagingCustomer.NewRepository(s.DB)
	scService := stagingCustomer.NewService(scRepo)
	scHandler := stagingCustomer.NewHandler(scService)

	s.Router.GET("/", scHandler.GetStagingCustomer)
}
