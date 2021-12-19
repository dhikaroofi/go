package controller

import (
	"errors"
	"github.com/dhikaroofi/go/app/helper"
	"github.com/dhikaroofi/go/app/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetDailyReportMerchant(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	getRepo, status, err := service.InitReportMerchantService(db, r)
	if err != nil {
		helper.RespondError(w, status, err, nil)
		return
	}
	vars := mux.Vars(r)
	merchantID, exist := vars["merchant_id"]
	if !exist {
		helper.RespondError(w, http.StatusBadRequest, errors.New("Invalid merchant id"), nil)
		return
	}
	merchantIDInt, err := strconv.ParseInt(merchantID, 10, 64)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, errors.New("Invalid merchant id"), nil)
		return
	}
	response, err := getRepo.GetDailyOmzet(merchantIDInt)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	helper.RespondSuccess(w, http.StatusOK, "Succes", response)
	return
}

func GetDailyReportOutlet(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	getRepo, status, err := service.InitReportOutletService(db, r)
	if err != nil {
		helper.RespondError(w, status, err, nil)
		return
	}
	vars := mux.Vars(r)
	outletID, exist := vars["outlet_id"]
	if !exist {
		helper.RespondError(w, http.StatusBadRequest, errors.New("Invalid merchant id"), nil)
		return
	}
	outletIDInt, err := strconv.ParseInt(outletID, 10, 64)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, errors.New("Invalid merchant id"), nil)
		return
	}
	response, err := getRepo.GetDailyReport(outletIDInt)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	helper.RespondSuccess(w, http.StatusOK, "Succes", response)
	return
}
