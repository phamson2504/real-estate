package response

type UserResponse struct {
	Id        int           `json:"id"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Role      string        `json:"role"`
	AvatarUrl string        `json:"avatar"`
	Contact   string        `json:"contact"`
	Agent     AgentResponse `json:"agent"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
