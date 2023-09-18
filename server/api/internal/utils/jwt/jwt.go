package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthService struct {
	AccessSecret string
	AccessExpire int64
}

func NewAuthService(secret string, expire int64) *AuthService {
	return &AuthService{
		AccessSecret: secret,
		AccessExpire: expire,
	}
}

func (s *AuthService) GetJwtToken(payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + s.AccessExpire
	claims["iat"] = time.Now().Unix()
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(s.AccessSecret))
}
