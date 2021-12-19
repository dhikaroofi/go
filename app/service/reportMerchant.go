package service

import (
	"errors"
	"github.com/dhikaroofi/go/app/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ReportMerchantService struct {
	Request struct {
		Parameter struct {
			StartDate  string
			EndDate    string
			Limit      int
			Page       int
			Offset     int
			MerchantID int64
		}
	}
	Response struct {
		pagination
	}
	baseService
}

func InitReportMerchantService(db *gorm.DB, r *http.Request) (service ReportMerchantService, statusCode int, err error) {
	service.r = r
	service.db = db
	service.getQueryParameter()
	if err := service.DecodeToken(); err != nil {
		return service, http.StatusUnauthorized, err
	}
	return service, http.StatusOK, nil
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
	keys, ok = service.r.URL.Query()["start_date"]
	if ok {
		service.Request.Parameter.StartDate = keys[0]
	} else {
		service.Request.Parameter.StartDate = "2021-11-01"
	}
	keys, ok = service.r.URL.Query()["end_date"]
	if ok {
		service.Request.Parameter.EndDate = keys[0]
	} else {
		service.Request.Parameter.EndDate = "2021-11-30"
	}

	service.Request.Parameter.Offset = service.Request.Parameter.Limit * service.Request.Parameter.Page
}

func (service *ReportMerchantService) GetDailyOmzet(merchantID int64) (interface{}, error) {
	if err := service.validate(merchantID); err != nil {
		return nil, err
	}
	var model []responseReport
	startDate := string(service.Request.Parameter.StartDate[0:8])
	sql := " SELECT DATE(calender.date) 'date',trx.omzet,trx.merchant_name  " +
		" FROM ( " +
		"    SELECT " +
		"        FROM_UNIXTIME(UNIX_TIMESTAMP(CONCAT(?,n)),'%Y-%m-%d') as date  " +
		"    FROM ( " +
		"            SELECT (((b4.0 << 1 | b3.0) << 1 | b2.0) << 1 | b1.0) << 1 | b0.0 as n " +
		"                    FROM  (SELECT 0 union all SELECT 1) as b0," +
		"                        (SELECT 0 union all SELECT 1) as b1," +
		"                        (SELECT 0 union all SELECT 1) as b2," +
		"                        (SELECT 0 union all SELECT 1) as b3," +
		"                        (SELECT 0 union all SELECT 1) as b4 ) t" +
		"            where n > 0 and n <= day(last_day(?)) ORDER BY date) calender" +
		" LEFT JOIN (" +
		"    SELECT " +
		"        SUM(trx.bill_total) 'omzet', " +
		"        mct.merchant_name 'merchant_name'," +
		"        DATE(trx.created_at) 'date' " +
		"    FROM Transactions trx " +
		"    INNER JOIN " +
		"        Merchants mct on mct.id=trx.merchant_id " +
		"    WHERE trx.merchant_id = ? AND trx.created_at " +
		"    BETWEEN ? AND ? GROUP BY date ORDER BY date) trx " +
		" ON calender.date=trx.date ORDER BY calender.date"

	execute := service.db.Raw(sql+" LIMIT ? OFFSET ? ", startDate,
		service.Request.Parameter.EndDate,
		merchantID,
		service.Request.Parameter.StartDate,
		service.Request.Parameter.EndDate,
		service.Request.Parameter.Limit,
		service.Request.Parameter.Offset).Scan(&model)
	if execute.Error != nil {
		return nil, execute.Error
	}
	service.Response.Page = service.Request.Parameter.Page
	service.Response.Limit = service.Request.Parameter.Limit
	service.Response.Lists = model
	return service.Response, nil
}

func (service ReportMerchantService) validate(merchantID int64) error {
	var count int64
	service.db.Model(&model.Merchants{}).Where("id = ? AND user_id = ?", merchantID, service.User.ID).Count(&count)
	if count < 1 {
		return errors.New("data tidak dapat diakses")
	}
	return nil
}
