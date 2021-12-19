package controller

import (
	"github.com/dhikaroofi/go/app/helper"
	"github.com/dhikaroofi/go/app/service"
	"gorm.io/gorm"
	"net/http"
)

func AuthLogin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	getRepo, err := service.InitLoginService(db, r)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	response, err := getRepo.Login()
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, err, nil)
		return
	}
	//if err != nil
	helper.RespondSuccess(w, http.StatusOK, "Succes", response)
	return
}

// func GetToken(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	authService := model.Auth{}
// 	if err := helper.DecodeJson(r,&authService.Request);err != nil {
// 		helper.RespondJSONError(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	data,err := authService.Login(db)
// 	if err != nil {
// 		helper.RespondJSONError(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	helper.RespondJSON(w, "Succes",http.StatusOK, data)
// }
