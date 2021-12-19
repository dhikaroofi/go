package service

import (
	"errors"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ReportMerchantService struct {
	Request struct {
		Parameter struct {
			Limit      int
			Page       int
			Offset     int
			MerchantID int64
		}
	}
	Response struct {
		responseReport
	}
	baseService
}

type responseReport struct {
	MerchantName string    `json:"merchant_name,omitempty"`
	OutletName   string    `json:"outlet_name,omitempty"`
	Omzet        float64   `json:"omzet,omitempty"`
	Date         time.Time `json:"date,omitempty"`
}

func InitReportMerchantService(db *gorm.DB, r *http.Request) (ReportMerchantService, error) {
	service := ReportMerchantService{}
	service.r = r
	service.db = db
	service.getQueryParameter()
	if err := service.DecodeToken(); err != nil {
		return service, err
	}
	return service, nil
}

func (service *ReportMerchantService) getQueryParameter() {
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

func (service *ReportMerchantService) GetDailyOmzet(merchantID int64) (interface{}, error) {
	if err := service.validate(merchantID); err != nil {
		return nil, err
	}
	var model []responseReport
	execute := service.db.Raw("SELECT "+
		"sum(trx.bill_total) 'omzet', "+
		"mct.merchant_name 'merchant_name', "+
		"DATE(trx.created_at) 'date' "+
		"FROM Transactions trx "+
		"INNER JOIN Merchants mct on mct.id=trx.merchant_id "+
		"WHERE trx.merchant_id = ? GROUP BY date LIMIT ? OFFSET ? ", merchantID, service.Request.Parameter.Limit, service.Request.Parameter.Offset).Scan(&model)
	if execute.Error != nil {
		return nil, execute.Error
	}
	return model, nil
}

func (service ReportMerchantService) validate(merchantID int64) error {
	var count int64
	service.db.Model(&model.Merchants{}).Where("id = ? AND user_id = ?", merchantID, service.User.ID).Count(&count)
	if count < 1 {
		return errors.New("data tidak dapat diakses")
	}
	return nil
}