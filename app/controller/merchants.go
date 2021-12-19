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

func GetListMerchants(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	getRepo, statusCode, err := service.InitMerchantService(db, r)
	if err != nil {
		helper.RespondError(w, statusCode, err, nil)
		return
	}
	response, err := getRepo.GetList()
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	helper.RespondSuccess(w, http.StatusOK, "Succes", response)
	return
}

func GetOutletList(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	getRepo, statusCode, err := service.InitMerchantService(db, r)
	if err != nil {
		helper.RespondError(w, statusCode, err, nil)
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
	response, err := getRepo.GetOutletList(merchantIDInt)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	helper.RespondSuccess(w, http.StatusOK, "Succes", response)
	return
}
