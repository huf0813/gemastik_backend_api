package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserID     int64  `json:"user_id"`
}

func NewAuthMiddleware(security string) (middleware.JWTConfig, error) {
	secret := []byte(security)
	isAuthorized := middleware.JWTConfig{
		SigningKey: secret,
	}
	return isAuthorized, nil
}

func NewToken(security, name string, userID int64) (string, error) {
	claims := &CustomClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:  "gudity_service",
			Subject: name,
		},
		UserID:     userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(security)
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func NewExtractToken(security, tokenFromUser string) (CustomClaim, error) {
	tokenFromUserWithoutBearer := strings.Replace(tokenFromUser, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenFromUserWithoutBearer, func(token *jwt.Token) (interface{}, error) {
		result := []byte(security)
		return result, nil
	})
	if err != nil {
		return CustomClaim{}, err
	}

	tokenClaims := token.Claims
	claims := tokenClaims.(jwt.MapClaims)
	customToken := CustomClaim{
		UserID:     int64(claims["user_id"].(float64)),
	}

	return customToken, nil
}
