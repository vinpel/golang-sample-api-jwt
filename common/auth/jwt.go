package auth

import (
	"errors"
	"fmt"

	"strings"
	"time"

	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username  string `json:"username" example:"John Smith"`
	Usage     string `json:"usage" example:"access token"`
	Empreinte string `json:"empreinte" example:"dazdza"`
	jwt.StandardClaims
}
type JWTTokens struct {
	Refresh string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZpbmNlbnQiLCJ1c2FnZSI6InJlZnJlc2ggdG9rZW4iLCJlbXByZWludGUiOiJTSnVod3BjSzJja3YxMSIsImV4cCI6MTY2ODI0NjQ3MX0.-Ga9cMnfN_6kNnYTsQCcF3FEbjXoaRmR4z4snJZYVbo"`
	Access  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZpbmNlbnQiLCJ1c2FnZSI6ImFjY2VzcyB0b2tlbiIsImVtcHJlaW50ZSI6IlNKdWh3cGNLMmNrdjExIiwiZXhwIjoxNjY2OTU0MDcxfQ.2RZfxWOy6_zyBJ9oFtsxr4Y5Zuh79qhDxhEJppbey3o"`
}

type TokenUsage int

const (
	Refresh TokenUsage = iota
	Access
)

func (s TokenUsage) String() string {
	switch s {
	case Refresh:
		return "refresh token"
	case Access:
		return "access token"
	}
	return "unknown"
}

func GenerateJWT(username string) (tokens JWTTokens, err error) {

	var user models.User
	record := database.Instance.Where("username = ?", username).First(&user)
	if record.Error != nil {
		err = errors.New("username not found")
		return
	}

	expirationTime := time.Now().Add(15 * time.Hour * 24)

	claims := &JWTClaim{
		Username:  username,
		Usage:     Refresh.String(),
		Empreinte: user.Empreinte,
	}
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokens.Refresh, _ = token.SignedString(jwtKey)

	expirationTime = time.Now().Add(1 * time.Hour)
	claims.Usage = Access.String()
	claims.ExpiresAt = expirationTime.Unix()

	tokens.Access, err = token.SignedString(jwtKey)

	return
}

func ValidateToken(authorizationString string, usage TokenUsage) (username string, err error) {
	// Split by space Bearer XxXXXxX
	parts := strings.SplitN(authorizationString, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		err = errors.New("format of request string need to be : Bearer <token>")
		return
	}

	token, err := jwt.ParseWithClaims(
		parts[1],
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	var user models.User
	record := database.Instance.Where("username = ?", claims.Username).First(&user)
	if record.Error != nil {
		err = fmt.Errorf("username '%v' not found", claims.Username)
		return
	}
	if user.Empreinte != claims.Empreinte {
		err = fmt.Errorf("password empreinte has changed, please request a new token from user data")
		return

	}
	if usage.String() != claims.Usage {
		err = fmt.Errorf("wrong token used, %v needed, %v used", usage, claims.Usage)
		return
	}
	username = claims.Username
	return
}
