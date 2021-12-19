package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dhikaroofi/go/app/model"
	"log"
	"os"
)

type AuthService struct {
	baseService
}

func InitAuthService() AuthService {
	return AuthService{}
}

func (service *AuthService) GenerateJWT(payload interface{}) error {
	var mySigningKey = []byte(os.Getenv("APP_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = payload.(model.Users).ID
	claims["user_name"] = payload.(model.Users).UserName
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Something Went Wrong:" + err.Error())
		return err
	}
	service.Token = tokenString
	return nil
}

func (service AuthService) ValidateJWT() (jwt.MapClaims, error) {
	token, err := jwt.Parse(service.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(os.Getenv("APP_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.New("Invalid token")
	}
	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("Invalid token:Expired")
		}
	}
	return nil, errors.New("Invalid token")
}
