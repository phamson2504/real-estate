package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/model"
	"real-estate-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PropertyController struct {
	PropertyService service.PropertyService
	FavorateService service.FavorateService
}

func NewPropertyController(
	propertyService service.PropertyService,
	favorateService service.FavorateService,
) *PropertyController {
	return &PropertyController{
		PropertyService: propertyService,
		FavorateService: favorateService,
	}
}

func (controller *PropertyController) Create(ctx *gin.Context) {
	log.Info().Msg("create property")

	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	usr, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to create upload directory"})
			return
		}
	}

	files, ok := form.File["images"]
	if !ok || len(files) == 0 {
		ctx.JSON(400, gin.H{"error": "No images found in the request"})
		return
	}

	var avatarURL []string
	for _, file := range files {
		filePath := filepath.Join(uploadDir, file.Filename)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}
		avatarURL = append(avatarURL, fmt.Sprintf("uploads/%s", file.Filename)) // Return relative URL
	}

	dataString := ctx.PostForm("data")
	var property request.PropertyData
	if err := json.Unmarshal([]byte(dataString), &property); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Xử lý trước khi khởi tạo struct
	maxPrice, err := strconv.ParseFloat(property.MaxPrice, 64)
	if err != nil {
		fmt.Println("Error convert MaxPrice:", err)
		return
	}

	minPrice, err := strconv.ParseFloat(property.MinPrice, 64)
	if err != nil {
		fmt.Println("Error convert MinPrice:", err)
		return
	}

	bedrooms, err := strconv.Atoi(property.Bedrooms)
	if err != nil {
		fmt.Println("Error convert Bedrooms:", err)
		return
	}

	bathrooms, err := strconv.Atoi(property.Bathrooms)
	if err != nil {
		fmt.Println("Error convert Bathrooms:", err)
		return
	}

	squareFeet, err := strconv.Atoi(property.SquareFeet)
	if err != nil {
		fmt.Println("Error convert SquareFeet:", err)
		return
	}

	propertyCreateRequest := request.PropertyCreateRequest{
		Title:       property.Title,
		Description: property.Description,
		MaxPrice:    maxPrice,
		MinPrice:    minPrice,
		Location:    property.Location,
		Bedrooms:    bedrooms,
		Bathrooms:   bathrooms,
		SquareFeet:  squareFeet,
		ImageURLs:   avatarURL,
	}

	errCreateProperty := controller.PropertyService.Create(usr, propertyCreateRequest)
	if errCreateProperty != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property", "details": errCreateProperty.Error()})
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

func (controller *PropertyController) Update(ctx *gin.Context) {
	log.Info().Msg("update property")
	propertyUpdateRequest := request.PropertyUpdateRequest{}
	err := ctx.ShouldBindJSON(&propertyUpdateRequest)
	helper.PanicIfError(err)

	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	usr, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	errUpdateProperty := controller.PropertyService.Update(usr, propertyUpdateRequest)
	if errUpdateProperty != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property", "details": errUpdateProperty.Error()})
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

func (controller *PropertyController) FindByPages(ctx *gin.Context) {
	log.Info().Msg("findByPages property")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	propertiesResp, totalProperties, totalPages := controller.PropertyService.FindByPages(page, limit)

	ctx.JSON(http.StatusOK, gin.H{
		"properties":      propertiesResp,
		"totalProperties": totalProperties,
		"totalPages":      totalPages,
		"currentPage":     page,
	})
}

func (controller *PropertyController) GetAllProperties(ctx *gin.Context) {
	log.Info().Msg("get All Properties")
	propertiesResp := controller.PropertyService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{
		"properties": propertiesResp,
	})
}

func (controller *PropertyController) GetPropertyById(ctx *gin.Context) {
	log.Info().Msg("Get Property By Ids")

	propertyId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(propertyId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}
	propertyResp := controller.PropertyService.FindById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"property": propertyResp,
	})
}

func (controller *PropertyController) GetPropertyByAgentId(ctx *gin.Context) {
	log.Info().Msg("Get Property By Agent Id")
	userId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}
	propertyResp := controller.PropertyService.FindByAgentId(id)

	ctx.JSON(http.StatusOK, gin.H{
		"properties": propertyResp,
	})
}
func (controller *PropertyController) GetPropertyFavoreat(ctx *gin.Context) {
	log.Info().Msg("Get Property Favoreat")
	propertyId := ctx.DefaultQuery("propertyId", "0")
	properId, err := strconv.Atoi(propertyId)
	if err != nil || properId <= 0 {
		helper.PanicIfError(err)
	}

	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	usr, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	favorReq := request.FavorateCreateRequest{
		UserId:     usr.Id,
		PropertyId: properId,
	}

	controller.FavorateService.Create(favorReq)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PropertyController) CheckPropertyFavoreat(ctx *gin.Context) {
	log.Info().Msg("Get Property Favoreat")
	propertyId := ctx.DefaultQuery("propertyId", "0")
	properId, err := strconv.Atoi(propertyId)
	if err != nil || properId <= 0 {
		helper.PanicIfError(err)
	}

	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	usr, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	favorReq := request.FavorateCreateRequest{
		UserId:     usr.Id,
		PropertyId: properId,
	}

	check := controller.FavorateService.CheckPropertyFavorite(favorReq)

	ctx.JSON(http.StatusOK, gin.H{
		"checkIsFavorite": check,
	})
}
func (controller *PropertyController) PropertyFavoriteByUserId(ctx *gin.Context) {
	log.Info().Msg("Get Property Favoreat")
	user, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	usr, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	properties := controller.FavorateService.PropertyFavoriteByUserId(usr.Id)

	ctx.JSON(http.StatusOK, gin.H{
		"properties": properties,
	})
}
