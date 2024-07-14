package config

import (
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/config/db_config"
	"log"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDBConfig()

	log.Println("Configurations loaded successfully.")
}
