package dto

type UserRequest struct {
	Name     string `json:"name" example:"John Doe"`
	Username string `json:"username" example:"john"`
	Email    string `json:"email" example:"demo@http://example.com/"`
	Password string `json:"password" example:"demo"`
}
