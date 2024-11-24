package response

type PropertyResponse struct {
	Id          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	MaxPrice    float64         `json:"maxPrice"`
	MinPrice    float64         `json:"minPrice"`
	Location    string          `json:"location"`
	Bedrooms    int             `json:"bedrooms"`
	Bathrooms   int             `json:"bathrooms"`
	SquareFeet  int             `json:"square_feet"`
	Agent       AgentResponse   `json:"agent"`
	Status      string          `json:"status"`
	Images      []ImageResponse `json:"image"`
}

type ImageResponse struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type AgentResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	AvatarAgent string `json:"avatarAgent"`
	Email       string `json:"email"`
}
type FavorateResponse struct {
	Id         int `json:"id"`
	UserId     int `json:"userId"`
	PropertyId int `json:"propertyId"`
}
