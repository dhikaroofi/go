package service

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ReportOutletService struct {
	Request struct {
		Parameter struct {
			Limit    int
			Page     int
			Offset   int
			OutletID int64
		}
	}
	Response struct {
		responseReport
	}
	baseService
}

func InitReportOutletService(db *gorm.DB, r *http.Request) (ReportOutletService, error) {
	service := ReportOutletService{}
	service.r = r
	service.db = db
	service.getQueryParameter()
	if err := service.DecodeToken(); err != nil {
		return service, err
	}
	return service, nil
}

func (service *ReportOutletService) getQueryParameter() {
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

func (service *ReportOutletService) GetDailyReport(outletID int64) (interface{}, error) {
	if err := service.validate(outletID); err != nil {
		return nil, err
	}
	var model []responseReport
	execute := service.db.Raw("SELECT "+
		"sum(trx.bill_total) 'omzet', "+
		"ot.merchant_name 'merchant_name',"+
		" ot.outlet_name 'outlet_name', "+
		"DATE(trx.created_at) 'date' "+
		"FROM Transactions trx "+
		"INNER JOIN (SELECT Outlets.id,Outlets.merchant_id,Outlets.outlet_name,Merchants.merchant_name "+
		"				FROM Outlets INNER JOIN Merchants ON Merchants.id=Outlets.merchant_id) ot "+
		"on ot.id=trx.outlet_id "+
		"WHERE trx.outlet_id = ? GROUP BY date LIMIT ? OFFSET ?; ", outletID, service.Request.Parameter.Limit, service.Request.Parameter.Offset).Scan(&model)
	if execute.Error != nil {
		return nil, execute.Error
	}
	return model, nil
}

func (service ReportOutletService) validate(outletid int64) error {
	var count int64
	execute := service.db.Raw("SELECT COUNT(*) count "+
		"FROM Outlets "+
		"INNER JOIN Merchants ON Merchants.id=Outlets.merchant_id "+
		"WHERE Merchants.user_id = ? AND Outlets.id = ?",
		service.User.ID, outletid).Scan(&count)
	if execute.Error != nil {
		return errors.New("data tidak dapat diakses")
	}
	if count < 1 {
		return errors.New("data tidak dapat diakses")
	}
	return nil
}
