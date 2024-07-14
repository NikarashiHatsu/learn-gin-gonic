package db_config

import "gin-gonic-gorm/config/db_utils"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1" // http://localhost
var DB_PORT = "3306"
var DB_NAME = "go_gin_gonic"
var DB_USER = "root"
var DB_PASS = "root"

func InitDBConfig() {
	db_utils.SetEnvIfNotEmpty("DB_DRIVER", &DB_DRIVER)
	db_utils.SetEnvIfNotEmpty("DB_HOST", &DB_HOST)
	db_utils.SetEnvIfNotEmpty("DB_PORT", &DB_PORT)
	db_utils.SetEnvIfNotEmpty("DB_NAME", &DB_NAME)
	db_utils.SetEnvIfNotEmpty("DB_USER", &DB_USER)
	db_utils.SetEnvIfNotEmpty("DB_PASS", &DB_PASS)
}
