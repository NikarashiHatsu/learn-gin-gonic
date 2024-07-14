package bootstrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// Load the .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error occured while loading .env file")
	}

	// Init config
	config.InitConfig()

	// Connect to the database
	database.ConnectDatabase()

	// Bootstrap the application
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
