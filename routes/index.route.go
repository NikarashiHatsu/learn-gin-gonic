package routes

import (
	"gin-gonic-gorm/controllers/book_controller"
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	// ROUTE USER
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/:id", user_controller.GetById)
	route.POST("/user", user_controller.Store)
	route.PUT("/user/:id", user_controller.UpdateById)
	route.DELETE("/user/:id", user_controller.DeleteById)

	// ROUTE BOOK
	route.GET("/book", book_controller.GetAllBook)
}
