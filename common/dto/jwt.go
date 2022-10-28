package dto

type TokenRequest struct {
	Email    string `json:"email"  example:"demo@example.net"`
	Password string `json:"password"  example:"demo"`
}
