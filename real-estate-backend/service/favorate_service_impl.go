package service

import (
	"fmt"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/model"
	"real-estate-backend/repository"
)

type FavorateServiceImpl struct {
	FavorateRepsitory  repository.FavorateRepsitory
	AgentRepository    repository.AgentRepository
	PropertyRepository repository.PropertyRepository
	ImageRepository    repository.ImageRepository
}

func NewFavorateServiceImpl(
	favorateRepsitory repository.FavorateRepsitory,
	agentRepository repository.AgentRepository,
	propertyRepository repository.PropertyRepository,
	imageRepository repository.ImageRepository,
) FavorateService {
	return &FavorateServiceImpl{
		FavorateRepsitory:  favorateRepsitory,
		AgentRepository:    agentRepository,
		PropertyRepository: propertyRepository,
		ImageRepository:    imageRepository,
	}
}

// Create implements FavorateService.
func (f *FavorateServiceImpl) Create(favorateReq request.FavorateCreateRequest) {
	favorate := model.Favorite{
		PropertyId: favorateReq.PropertyId,
		UserId:     favorateReq.UserId,
	}
	f.FavorateRepsitory.Save(favorate)
}

// CheckPropertyFavorite implements FavorateService.
func (f *FavorateServiceImpl) CheckPropertyFavorite(favorateReq request.FavorateCreateRequest) bool {
	favorate := model.Favorite{
		PropertyId: favorateReq.PropertyId,
		UserId:     favorateReq.UserId,
	}

	return f.FavorateRepsitory.CheckPropertyFavorite(favorate)
}

// PropertyFavoriteByUserId implements FavorateService.
func (f *FavorateServiceImpl) PropertyFavoriteByUserId(userId int) []response.PropertyResponse {
	favarates := f.FavorateRepsitory.GetPropertyFavoriteByUserId(userId)
	fmt.Printf("123333333333333fdafada%d", userId)
	baseURL := "http://localhost:8080"
	properties := []response.PropertyResponse{}
	for _, favarate := range favarates {
		property := f.PropertyRepository.FindById(favarate.PropertyId)
		agent := f.AgentRepository.FindByAgentId(property.AgentId)
		var images = f.ImageRepository.FindByPropertyID(property.Id)

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
		var status string
		if property.Status == 0 {
			status = "spending"
		} else if property.Status == 1 {
			status = "verified"
		} else {
			status = "declined"
		}
		propertyResp := response.PropertyResponse{
			Id:          property.Id,
			Title:       property.Title,
			Description: property.Description,
			MaxPrice:    property.MaxPrice,
			MinPrice:    property.MaxPrice,
			Location:    property.Location,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			Status:      status,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}
		properties = append(properties, propertyResp)
	}
	return properties
}
