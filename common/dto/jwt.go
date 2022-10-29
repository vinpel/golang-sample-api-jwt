package dto

type TokenRequest struct {
	Email    string `json:"email"  example:"demo@example.com"`
	Password string `json:"password"  example:"demo"`
}
