package service

import (
	"errors"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type MerchantService struct {
	Request struct {
		Parameter struct {
			Limit      int   `json:"limit,omitempty"`
			Page       int   `json:"page,omitempty"`
			Offset     int   `json:"offset,omitempty"`
			MerchantID int64 `json:"merchant_id,omitempty"`
		}
	}
	Response struct {
		pagination
	}
	baseService
}

func InitMerchantService(db *gorm.DB, r *http.Request) (service MerchantService, statusCode int, err error) {
	service.r = r
	service.db = db
	service.getQueryParameter()
	if err := service.DecodeToken(); err != nil {
		return service, http.StatusUnauthorized, err
	}
	return service, http.StatusOK, nil
}

func (service *MerchantService) getQueryParameter() {
	service.Request.Parameter.Limit = 10
	service.Request.Parameter.Page = 0
	keys, ok := service.r.URL.Query()["limit"]
	if ok {
		limit, err := strconv.Atoi(keys[0])
		if err == nil {
			service.Request.Parameter.Limit = limit
		}
	}
	keys, ok = service.r.URL.Query()["page"]
	if ok {
		page, err := strconv.Atoi(keys[0])
		if err == nil {
			service.Request.Parameter.Page = page
		}
	}

	service.Request.Parameter.Offset = service.Request.Parameter.Limit * service.Request.Parameter.Page
}

func (service *MerchantService) GetList() (interface{}, error) {
	var model []model.Merchants
	execute := service.db.Where("user_id", service.User.ID).Find(&model).Limit(service.Request.Parameter.Limit).Offset(service.Request.Parameter.Offset).Order("created_at")
	if execute.Error != nil {
		return nil, execute.Error
	}
	service.Response.Page = service.Request.Parameter.Page
	service.Response.Limit = service.Request.Parameter.Limit
	service.Response.Lists = model
	return service.Response, nil
}

func (service *MerchantService) GetOutletList(merchantID int64) (interface{}, error) {
	if err := service.validate(merchantID); err != nil {
		return nil, err
	}
	var model []model.Outlets
	execute := service.db.Where("merchant_id", merchantID).Find(&model).Limit(service.Request.Parameter.Limit).Offset(service.Request.Parameter.Offset).Order("created_at")
	if execute.Error != nil {
		return nil, execute.Error
	}
	service.Response.Page = service.Request.Parameter.Page
	service.Response.Limit = service.Request.Parameter.Limit
	service.Response.Lists = model
	return service.Response, nil
}

func (service MerchantService) validate(merchantID int64) error {
	var count int64
	service.db.Model(&model.Merchants{}).Where("id = ? AND user_id = ?", merchantID, service.User.ID).Count(&count)
	if count < 1 {
		return errors.New("data tidak dapat diakses")
	}
	return nil
}
