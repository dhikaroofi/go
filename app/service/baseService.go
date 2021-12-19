package service

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type baseService struct {
	r       *http.Request
	db      *gorm.DB
	User    model.Users
	JwtData jwt.MapClaims
	Token   string
}

type pagination struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Lists interface{} `json:"lists,omitempty"`
}

type responseReport struct {
	MerchantName string    `json:"merchant_name,omitempty"`
	OutletName   string    `json:"outlet_name,omitempty"`
	Omzet        float64   `json:"omzet,omitempty"`
	Date         time.Time `json:"date,omitempty"`
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
