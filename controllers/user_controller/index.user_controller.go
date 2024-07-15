package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
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

	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})

		return
	}

	if user.ID == nil {
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
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error

	if errDb != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error occurred",
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User stored.",
		"data":    user,
	})
}

func UpdateById(ctx *gin.Context) {

}

func DeleteById(ctx *gin.Context) {

}
