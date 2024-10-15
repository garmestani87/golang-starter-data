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


// @Summary save user
// @Description save user
// @Accept  json
// @Produce  json
// @param UserDto	body	dto.UserDto	true	"User Dto"
// @Success 200 {object} common.ResponseModel[dto.UserDto, dto.UserDto]
// @Failure 500 {object} common.ResponseModel[dto.UserDto, dto.UserDto]
// @Router /api/v1/user/ [post]
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

// @Summary find user by id
// @Description find user by id
// @Accept  json
// @Produce  json
// @Param   id     path    int     true     "id"
// @Success 200 {object} dto.UserDto
// @Failure 500 {object} string
// @Router /api/v1/user/{id} [get]
func FindUserById(ctx *gin.Context) {
	user := entity.NewUserEntity()
	res := db.GetDatabase().First(&user, ctx.Param("id"))
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": res.Error.Error()})
		return
	}
	ctx.JSON(200, mapper.MapToDto(user))
}
