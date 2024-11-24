package controller

import (
	"net/http"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("create user")
	userCreateRequest := request.UserCreateRequest{}
	err := ctx.ShouldBindJSON(&userCreateRequest)
	helper.PanicIfError(err)

	controller.UserService.Create(userCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tags")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("update tags")
	userUpdateRequest := request.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(&userUpdateRequest)
	helper.PanicIfError(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)
	userUpdateRequest.Id = id

	controller.UserService.Update(userUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	result := controller.UserService.FindById(id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	result := controller.UserService.FindAll()

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
