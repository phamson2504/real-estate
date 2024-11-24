package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/model"
	"real-estate-backend/repository"
	"real-estate-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
	agentRepo             repository.AgentRepository
}

func NewAuthenticationController(service service.AuthenticationService, agentRepo repository.AgentRepository) *AuthenticationController {
	return &AuthenticationController{authenticationService: service, agentRepo: agentRepo}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	log.Info().Msg("login tags")
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.PanicIfError(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	log.Info().Msg("register tags")
	createUsersRequest := request.UserCreateRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.PanicIfError(err)

	errRegister := controller.authenticationService.Register(createUsersRequest)
	if errRegister != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: errRegister.Error(),
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) CurrentUser(ctx *gin.Context) {
	log.Info().Msg("currentUser tags")
	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	currentUser, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	agent, err := controller.agentRepo.FindByUserId(currentUser.Id)

	var agentResp = response.AgentResponse{}
	var avatarUrl = ""
	if err == nil {
		agentResp = response.AgentResponse{
			Id:      agent.Id,
			Contact: agent.ContactNumber,
			Name:    agent.AgencyName,
		}

		if agent.AvatarAgent != "" {
			baseURL := "http://localhost:8080"
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}
	}

	userResp := response.UserResponse{
		Id:        currentUser.Id,
		Username:  currentUser.Username,
		Email:     currentUser.Email,
		Role:      currentUser.Role,
		AvatarUrl: avatarUrl,
		Agent:     agentResp,
	}

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully geted user!",
		Data:    userResp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) UpdateProfile(ctx *gin.Context) {
	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	currentUser, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	agentId, _ := strconv.Atoi(ctx.PostForm("agentId"))
	contact := ctx.PostForm("contact")
	email := ctx.PostForm("email")
	agentName := ctx.PostForm("agentName")

	// get file from FormData
	file, err := ctx.FormFile("avatar")
	var avatarURL string
	if err == nil {
		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.MkdirAll(uploadDir, os.ModePerm)
		}

		filePath := filepath.Join(uploadDir, file.Filename)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
			return
		}
		avatarURL = filePath
	}

	agent := model.Agent{
		Id:            agentId,
		AgencyName:    agentName,
		ContactNumber: contact,
		AvatarAgent:   avatarURL,
		UserId:        currentUser.Id,
	}

	controller.authenticationService.UpdateProfile(agent, email, *currentUser)

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Profile updated successfully",
		"contact":   contact,
		"email":     email,
		"agentName": agentName,
		"avatarUrl": avatarURL,
	})
}
