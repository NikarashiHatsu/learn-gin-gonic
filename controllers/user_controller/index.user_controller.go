package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)

	err := database.DB.Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error.",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Users fetched.",
		"data":    users,
	})
}
