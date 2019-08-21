package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"../auth"
	"../helper"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
)

const (
	db         = "Dormnet"
	collection = "easyShift"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	// Decode Json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}

	//check user is exist or not
	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: 42200, Msg: "該帳號已經存在囉QQ~"})
	} else {
		// Register new user

		// Hash password
		hasher := md5.New()
		hasher.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(hasher.Sum(nil))

		//  change roles if you want
		user.Roles = []string{"editor"}

		// Insert to database
		err = models.Insert(db, collection, user)
		if err != nil {
			helper.ResponseWithJson(w, http.StatusInternalServerError,
				helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		}

		fmt.Println("[INFO] " + user.UserName + " register success!")
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: 20000, Msg: "註冊成功"})
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	// Decode Json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}

	var result models.User
	err = models.FindOne(db, collection, bson.M{"username": user.UserName}, nil, &result)
	if err != nil {
		if err.Error() == "not found" {
			helper.ResponseWithJson(w, http.StatusOK,
				helper.Response{Code: 42200, Msg: "帳號或密碼錯誤"})
		} else {
			helper.ResponseWithJson(w, http.StatusInternalServerError,
				helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		}
		return
	}

	// Check Password
	hasher := md5.New()
	hasher.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hasher.Sum(nil))
	if user.Password != result.Password {
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: 42200, Msg: "帳號或密碼錯誤"})
		return
	}

	// Validate Success~
	// Generate JWT token
	token, _ := auth.GenerateToken(&user)
	fmt.Println("[INFO] " + user.UserName + " login success!")
	helper.ResponseWithJson(w, http.StatusOK,
		helper.Response{Code: 20000, Msg: user.UserName + " 登入成功~", Data: models.JwtToken{Token: token}})

}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	// Parse Token Information
	tokenString := r.Header.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// failed
			helper.ResponseWithJson(w, http.StatusUnauthorized,
				helper.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
			return nil, fmt.Errorf("not authorization")
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var result models.User
		err := models.FindOne(db, collection, bson.M{"username": claims["username"].(string)}, nil, &result)
		if err != nil {
			helper.ResponseWithJson(w, http.StatusInternalServerError,
				helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		} else {
			//success
			helper.ResponseWithJson(w, http.StatusOK,
				helper.Response{Code: 20000, Msg: "取得用戶資訊0 -0~", Data: models.UserInfo{Roles: result.Roles, Name: result.UserName}})
			return
		}

	} else {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		return
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	helper.ResponseWithJson(w, http.StatusOK,
		helper.Response{Code: 20000, Msg: "success"})
}
