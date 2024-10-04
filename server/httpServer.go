package server

import (
	"database/sql"
	"log"
	"hotel-booking/controllers"
	"hotel-booking/repositories"
	"hotel-booking/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	adminController *controllers.adminController
	// resultsController *controllers.ResultsController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	adminRepository := repositories.NewAdminRepository(dbHandler)
	adminService := services.NewAdminService(adminRepository)
	adminController := controllers.NewAdminController(adminService)

	router := gin.Default()

	router.POST("/admin", adminController.LoginHandler)
	

	return HttpServer{
		config:          config,
		router:          router,
		adminController: adminController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}