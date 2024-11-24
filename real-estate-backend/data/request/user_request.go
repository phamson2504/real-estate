package request

type UserCreateRequest struct {
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Password  string `validate:"required,min=8" json:"password"`
	Email     string `validate:"required,email,min=10,max=100" json:"email"`
	Role      string
	AvatarUrl string
}

type LoginRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserUpdateRequest struct {
	Id        int
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Password  string `validate:"required,min=8" json:"password"`
	Email     string `validate:"required,email,min=10,max=100" json:"email"`
	Role      string
	AvatarUrl string
}
type AgentRequest struct {
	Id          int
	Name        string
	Contact     string
	AvatarAgent string
}
