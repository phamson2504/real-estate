package request

type PropertyCreateRequest struct {
	Title       string
	Description string
	MaxPrice    float64
	MinPrice    float64
	Location    string
	Bedrooms    int
	Bathrooms   int
	SquareFeet  int
	ImageURLs   []string
}

type PropertyUpdateRequest struct {
	Id          int
	Title       string
	Description string
	MaxPrice    float64
	MinPrice    float64
	Location    string
	Bedrooms    int
	Bathrooms   int
	SquareFeet  int
	Image       []ImageRequest
}

type PropertyData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MaxPrice    string `json:"maxPrice"`
	MinPrice    string `json:"minPrice"`
	Location    string `json:"location"`
	Bedrooms    string `json:"bedrooms"`
	Bathrooms   string `json:"bathrooms"`
	SquareFeet  string `json:"squareFeet"`
}

type ImageRequest struct {
	Id  int
	URL string
}

type FavorateCreateRequest struct {
	UserId     int
	PropertyId int
}
