package api

import (
	"project_credit_sinarmas/backend/controllers/authentication"
	"project_credit_sinarmas/backend/controllers/report"
	"project_credit_sinarmas/backend/controllers/skalaAngsuran"
	"project_credit_sinarmas/backend/controllers/stagingCustomer"
	"project_credit_sinarmas/backend/controllers/transaction"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "DELETE", "PUT", "GET", "PATCH"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	authRepo := authentication.NewRepository(s.DB)
	authService := authentication.NewService(authRepo)
	authHandler := authentication.NewHandler(authService)
	s.Router.POST("/login", authHandler.Login)
	s.Router.GET("/user", authHandler.AuthenticateUser)
	s.Router.PATCH("/changePassword", authHandler.ChangePassword)

	scRepo := stagingCustomer.NewRepository(s.DB)
	scService := stagingCustomer.NewService(scRepo)
	scHandler := stagingCustomer.NewHandler(scService)
	s.Router.GET("/sc", scHandler.GetStagingCustomer)

	saRepo := skalaAngsuran.NewRepository(s.DB)
	saService := skalaAngsuran.NewService(saRepo)
	saHandler := skalaAngsuran.NewHandler(saService)
	s.Router.GET("/sa", saHandler.GenerateSkalaAngsuran)

	transactionRepo := transaction.NewRepository(s.DB)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService)
	s.Router.GET("/getTransaction", transactionHandler.GetTransaction)
	s.Router.GET("/getTransactionFilter", transactionHandler.GetTransactionFilter)
	s.Router.PATCH("/updateTransaction", transactionHandler.UpdateTransaction)

	reportRepo := report.NewRepository(s.DB)
	reportService := report.NewService(reportRepo)
	reportHandler := report.NewHandler(reportService)
	s.Router.GET("/getReport", reportHandler.GetReport)
	s.Router.GET("/getReportFilter", reportHandler.GetReportFilter)

}
