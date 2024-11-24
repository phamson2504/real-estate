package service

import (
	"errors"
	"fmt"
	"log"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/model"
	"real-estate-backend/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type PropertyServiceImpl struct {
	PropertyRepository repository.PropertyRepository
	AgentRepository    repository.AgentRepository
	ImageRepository    repository.ImageRepository
	UserRepository     repository.UserRepository
	Validate           *validator.Validate
}

func NewPropertyServiceImpl(propertyRepository repository.PropertyRepository,
	agentRepository repository.AgentRepository,
	imageRepository repository.ImageRepository,
	userRepository repository.UserRepository,
	validate *validator.Validate) PropertyService {

	return &PropertyServiceImpl{
		PropertyRepository: propertyRepository,
		AgentRepository:    agentRepository,
		ImageRepository:    imageRepository,
		UserRepository:     userRepository,
		Validate:           validate,
	}
}

// Create implements PropertyService.
func (p *PropertyServiceImpl) Create(user *model.User, property request.PropertyCreateRequest) error {
	agent, err := p.AgentRepository.FindByUserId(user.Id)
	if err != nil {
		return errors.New("agent not found")
	}

	newProperty := model.Property{
		Title:       property.Title,
		Location:    property.Location,
		Description: property.Description,
		MaxPrice:    property.MaxPrice,
		MinPrice:    property.MinPrice,
		AgentId:     agent.Id,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
		Status:      1,
		CreatedAt:   time.Now(),
	}
	p.PropertyRepository.Save(&newProperty)

	newImagesURLs := property.ImageURLs

	for _, value := range newImagesURLs {
		var image = model.Image{
			URL:        value,
			PropertyId: newProperty.Id,
		}
		p.ImageRepository.Save(image)
	}

	return nil
}

// Update implements PropertyService.
func (p *PropertyServiceImpl) Update(user *model.User, property request.PropertyUpdateRequest) error {
	agent, err := p.AgentRepository.FindByUserId(user.Id)
	if err != nil {
		return errors.New("agent not found")
	}

	updateProperty := model.Property{
		Id:          property.Id,
		Title:       property.Title,
		Location:    property.Location,
		Description: property.Description,
		MaxPrice:    property.MaxPrice,
		MinPrice:    property.MinPrice,
		AgentId:     agent.Id,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
	}
	images := p.ImageRepository.FindByPropertyID(property.Id)
	imageUpdate := property.Image

	updateMap := make(map[int]bool)
	for _, img := range imageUpdate {
		if img.Id == 0 {
			image := model.Image{
				Id:         img.Id,
				URL:        img.URL,
				PropertyId: property.Id,
			}
			p.ImageRepository.Save(image)
		} else {
			updateMap[img.Id] = true
		}
	}

	for _, img := range images {
		if !updateMap[img.Id] {
			p.ImageRepository.Delete(img.Id)
		}
	}

	p.PropertyRepository.Update(updateProperty)

	return nil
}

// FindAll implements PropertyService.
func (p *PropertyServiceImpl) FindByPages(page int, limit int) ([]response.PropertyResponse, int64, int) {
	offset := (page - 1) * limit
	properties := p.PropertyRepository.FindByOffset(offset, limit)
	totalProperties := p.PropertyRepository.TotalProperties()

	totalPages := int((totalProperties + int64(limit) - 1) / int64(limit))
	baseURL := "http://localhost:8080"

	var propertiesResp []response.PropertyResponse
	for _, value := range properties {
		var images = p.ImageRepository.FindByPropertyID(value.Id)
		var agent = p.AgentRepository.FindByAgentId(value.AgentId)

		// fmt.Printf(" %d dsssssssssssssssssssssss : %+v\n", value.Id, images)

		var status string
		if value.Status == 0 {
			status = "spending"
		} else if value.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			MaxPrice:    value.MaxPrice,
			MinPrice:    value.MinPrice,
			Location:    value.Location,
			Bedrooms:    value.Bedrooms,
			Bathrooms:   value.Bathrooms,
			SquareFeet:  value.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}

		propertiesResp = append(propertiesResp, propertyResp)
	}
	return propertiesResp, totalProperties, totalPages
}

func (p *PropertyServiceImpl) FindAll() []response.PropertyResponse {
	var propertiesResp []response.PropertyResponse
	properties := p.PropertyRepository.FindAll()
	baseURL := "http://localhost:8080"

	for _, value := range properties {
		var images = p.ImageRepository.FindByPropertyID(value.Id)
		var agent = p.AgentRepository.FindByAgentId(value.AgentId)
		var status string
		if value.Status == 0 {
			status = "spending"
		} else if value.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			MaxPrice:    value.MaxPrice,
			MinPrice:    value.MinPrice,
			Location:    value.Location,
			Bedrooms:    value.Bedrooms,
			Bathrooms:   value.Bathrooms,
			SquareFeet:  value.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}

		propertiesResp = append(propertiesResp, propertyResp)
	}
	return propertiesResp
}

// FindById implements PropertyService.
func (p *PropertyServiceImpl) FindById(propertyId int) response.PropertyResponse {
	baseURL := "http://localhost:8080"

	property := p.PropertyRepository.FindById(propertyId)

	var images = p.ImageRepository.FindByPropertyID(property.Id)
	var agent = p.AgentRepository.FindByAgentId(property.AgentId)
	userAgent, _ := p.UserRepository.FindById(agent.UserId)
	var status string
	if property.Status == 0 {
		status = "spending"
	} else if property.Status == 1 {
		status = "verified"
	} else {
		status = "declined"
	}

	var avatarUrl = ""
	if agent.AvatarAgent != "" {
		avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
	}

	agentResp := response.AgentResponse{
		Id:          userAgent.Id,
		Name:        agent.AgencyName,
		Contact:     agent.ContactNumber,
		AvatarAgent: avatarUrl,
		Email:       userAgent.Email,
	}

	propertyResp := response.PropertyResponse{
		Id:          property.Id,
		Title:       property.Title,
		Description: property.Description,
		MaxPrice:    property.MaxPrice,
		MinPrice:    property.MinPrice,
		Location:    property.Location,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
		Status:      status,
		Agent:       agentResp,
	}

	if len(images) > 0 {
		propertyResp.Images = make([]response.ImageResponse, 0, len(images))
		for _, image := range images {
			var imageUrl = fmt.Sprintf("%s/%s", baseURL, image.URL)
			propertyResp.Images = append(propertyResp.Images, response.ImageResponse{
				Id:  image.Id,
				Url: imageUrl,
			})
		}
	}

	return propertyResp
}

// FindByUserId implements PropertyService.
func (p *PropertyServiceImpl) FindByAgentId(agentId int) []response.PropertyResponse {
	var propertiesResp []response.PropertyResponse
	properties := p.PropertyRepository.FindByAgentId(agentId)
	baseURL := "http://localhost:8080"
	for _, value := range properties {
		var images = p.ImageRepository.FindByPropertyID(value.Id)
		var agent = p.AgentRepository.FindByAgentId(value.AgentId)
		var status string
		if value.Status == 0 {
			status = "spending"
		} else if value.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			MaxPrice:    value.MaxPrice,
			MinPrice:    value.MinPrice,
			Location:    value.Location,
			Bedrooms:    value.Bedrooms,
			Bathrooms:   value.Bathrooms,
			SquareFeet:  value.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}

		propertiesResp = append(propertiesResp, propertyResp)
	}
	return propertiesResp
}

// GetPropertiesBought implements PropertyService.
func (p *PropertyServiceImpl) GetPropertiesBought(userId int) []response.PropertyResponse {
	var propertiesResp []response.PropertyResponse
	properties := p.PropertyRepository.FindByBought(userId)
	baseURL := "http://localhost:8080"
	for _, value := range properties {
		var images = p.ImageRepository.FindByPropertyID(value.Id)
		var agent = p.AgentRepository.FindByAgentId(value.AgentId)
		var status string
		if value.Status == 0 {
			status = "spending"
		} else if value.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			MaxPrice:    value.MaxPrice,
			MinPrice:    value.MinPrice,
			Location:    value.Location,
			Bedrooms:    value.Bedrooms,
			Bathrooms:   value.Bathrooms,
			SquareFeet:  value.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}

		propertiesResp = append(propertiesResp, propertyResp)
	}
	return propertiesResp
}

// GetPropertiesSold implements PropertyService.
func (p *PropertyServiceImpl) GetPropertiesSold(userId int) []response.PropertyResponse {
	var propertiesResp []response.PropertyResponse
	properties := p.PropertyRepository.FindBySold(userId)
	baseURL := "http://localhost:8080"
	for _, value := range properties {
		var images = p.ImageRepository.FindByPropertyID(value.Id)
		var agent = p.AgentRepository.FindByAgentId(value.AgentId)
		var status string
		if value.Status == 0 {
			status = "spending"
		} else if value.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			MaxPrice:    value.MaxPrice,
			MinPrice:    value.MinPrice,
			Location:    value.Location,
			Bedrooms:    value.Bedrooms,
			Bathrooms:   value.Bathrooms,
			SquareFeet:  value.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}

		propertiesResp = append(propertiesResp, propertyResp)
	}
	return propertiesResp
}
