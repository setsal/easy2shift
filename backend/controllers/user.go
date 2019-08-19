package controllers

import (
	"encoding/json"
	"net/http"

	"../auth"
	"../helper"
	"../models"
	"github.com/globalsign/mgo/bson"
)

const (
	db         = "Dormnet"
	collection = "easyShift"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}

	//check user is exist or not
	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		helper.ResponseWithJson(w, http.StatusUnprocessableEntity,
			helper.Response{Code: http.StatusUnprocessableEntity, Msg: "User Alreay Exist!"})
	} else {
		//register new user
		err = models.Insert(db, collection, user)
		// hash := md5.Sum([]byte(user.UserName))
		// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err != nil {
			helper.ResponseWithJson(w, http.StatusInternalServerError,
				helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		}

		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: http.StatusOK, Msg: "register success!"})

	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}

	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		token, _ := auth.GenerateToken(&user)
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
	} else {
		helper.ResponseWithJson(w, http.StatusNotFound,
			helper.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
	}
}
