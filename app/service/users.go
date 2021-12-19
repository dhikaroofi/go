package service

import (
	"errors"
	"github.com/dhikaroofi/go/app/helper"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"net/http"
)

type UsersService struct {
	Request struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	Response struct {
		Users         model.Users `json:"users"`
		Authorization struct {
			Type  string `json:"type"`
			Token string `json:"token"`
		} `json:"authorization"`
		Pagination interface{}
	}
	baseService
}

func InitLoginService(db *gorm.DB, r *http.Request) (UsersService, error) {
	service := UsersService{}
	service.db = db
	if err := service.DecodeJson(r, &service.Request); err != nil {
		return service, errors.New("invalid payload")
	}
	return service, nil
}

func (service *UsersService) Login() (interface{}, error) {
	users := model.Users{}
	excute := service.db.Where("user_name = ?", service.Request.UserName).First(&users)
	if excute.Error != nil {
		return nil, excute.Error
	}
	if excute.RowsAffected < 1 {
		return nil, errors.New("username tidak ditemukan")
	}
	if !helper.ValidateMd5(users.Password, service.Request.Password) {
		return nil, errors.New("password salah")
	}
	if err := service.setToken(users); err != nil {
		return nil, err
	}
	service.Response.Users = users
	return service.Response, nil
}

func (service *UsersService) validateToken() error {
	return nil
}

func (service *UsersService) setToken(payload interface{}) error {
	token := InitAuthService()
	if err := token.GenerateJWT(payload); err != nil {
		return err
	}
	service.Response.Authorization.Type = "Bearer"
	service.Response.Authorization.Token = token.Token
	return nil
}
