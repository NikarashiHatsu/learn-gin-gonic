package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"

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

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	err := database.DB.Where("id = ?", id).Find(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "User not found.",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "User found.",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {

}

func UpdateById(ctx *gin.Context) {

}

func DeleteById(ctx *gin.Context) {

}
