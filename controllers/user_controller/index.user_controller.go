package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	httpErrorStatus, findUserByIdErr := findUserById(&id, user)

	if findUserByIdErr != "" {
		ctx.AbortWithStatusJSON(httpErrorStatus, gin.H{
			"message": findUserByIdErr,
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

	httpStatus, emailValidationCheckErr := checkIfEmailExists(&userReq.Email, nil)

	if emailValidationCheckErr != "" {
		ctx.AbortWithStatusJSON(httpStatus, gin.H{
			"message": emailValidationCheckErr,
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
	id := ctx.Param("id")

	user := new(responses.UserResponse)

	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	httpErrorStatus, findUserByIdErr := findUserById(&id, user)

	if findUserByIdErr != "" {
		ctx.AbortWithStatusJSON(httpErrorStatus, gin.H{
			"message": findUserByIdErr,
		})

		return
	}

	httpErrorStatus, emailValidationCheckErr := checkIfEmailExists(&userReq.Email, user.ID)

	if emailValidationCheckErr != "" {
		ctx.AbortWithStatusJSON(httpErrorStatus, gin.H{
			"message": emailValidationCheckErr,
		})

		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error

	if errUpdate != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "User updated.",
		"data":    user,
	})
}

func DeleteById(ctx *gin.Context) {

}

func findUserById(id *string, user *responses.UserResponse) (int, string) {
	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if err != nil {
		return http.StatusInternalServerError, "Internal server error"
	}

	if user.ID == nil {
		return http.StatusNotFound, "User not found."
	}

	return http.StatusOK, ""
}

func checkIfEmailExists(email *string, userId *int) (int, string) {
	userEmailExists := new(models.User)

	errUserEmailExists := database.DB.Table("users").Where("email = ?", email).First(&userEmailExists).Error

	if errUserEmailExists == gorm.ErrRecordNotFound {
		return http.StatusOK, ""
	}

	if errUserEmailExists != nil {
		return http.StatusInternalServerError, "Internal server error."
	}

	if userEmailExists.Email != nil && !userIsUpdatingItsData(*userId, *userEmailExists.ID) {
		return http.StatusBadRequest, "User with that email is already exists."
	}

	return http.StatusOK, ""
}

func userIsUpdatingItsData(userIdFromInitialFetch int, userIdFromMailValidation int) bool {
	return userIdFromInitialFetch == userIdFromMailValidation
}
