package dto
type TokenRequest struct {
	Email    string `json:"email"  example:"demo@exemple.net"`
	Password string `json:"password"  example:"demo"`
}
