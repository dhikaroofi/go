package service

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type baseService struct {
	r       *http.Request
	db      *gorm.DB
	User    model.Users
	JwtData jwt.MapClaims
	Token   string
}

func (baseService) DecodeJson(r *http.Request, payload interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func (service *baseService) DecodeToken() error {
	token := service.r.Header.Get("Authorization")
	if token == "" {
		return errors.New("Tidak diizinkan")
	}
	auth := AuthService{}
	auth.Token = token
	dataToken, err := auth.ValidateJWT()
	if err != nil {
		return err
	}
	service.JwtData = dataToken
	service.Token = token
	if err := service.getUser(); err != nil {
		return err
	}
	return nil
}

func (service *baseService) getUser() error {
	user := model.Users{}
	excute := service.db.Where("id = ?", service.JwtData["user_id"]).First(&user)
	if excute.Error != nil {
		return excute.Error
	}
	if excute.RowsAffected < 1 {
		return errors.New("username tidak ditemukan")
	}
	service.User = user
	return nil
}
