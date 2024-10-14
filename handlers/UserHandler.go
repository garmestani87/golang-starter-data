package handlers

import (
	"golang-starter-data/common"
	"golang-starter-data/db"
	"golang-starter-data/domain/dto"
	"golang-starter-data/domain/entity"
	"golang-starter-data/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUser(ctx *gin.Context) {
	userDto := dto.NewUserDto()
	err := ctx.ShouldBindBodyWithJSON(&userDto)
	response := common.NewResponseModel[dto.UserDto, int]()
	if err != nil {
		ex := common.NewExceptionModel("Error occurred when parsing request body !", common.EXCEPTION, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &ex)
		return
	}

	res := db.GetDatabase().Create(mapper.MapToEntity(userDto))

	if res.Error != nil {
		ex := common.NewExceptionModel("Error occurred when inserting user !", common.EXCEPTION, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &ex)
		return
	}
	response.Input = *userDto
	ctx.JSON(http.StatusOK, &response)
}

func FindUserById(ctx *gin.Context) {
	user := entity.NewUserEntity()
	res := db.GetDatabase().First(&user, ctx.Param("id"))
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": res.Error.Error()})
		return
	}
	ctx.JSON(200, user)
}
