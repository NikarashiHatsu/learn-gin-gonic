package app_config

import "gin-gonic-gorm/config/db_utils"

var PORT = ":8080"

func InitAppConfig() {
	db_utils.SetEnvIfNotEmpty("APP_PORT", &PORT)
}
