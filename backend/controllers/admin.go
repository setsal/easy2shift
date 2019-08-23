package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../helper"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
)

var (
	adminDao = models.Schedule{}
)

type Month struct {
	Month string `json:"month"`
}

func SetMonth(w http.ResponseWriter, r *http.Request) {

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
			var month Month
			err := json.NewDecoder(r.Body).Decode(&month)
			if err != nil {
				print("Error")
				return
			}

			// Insert data
			var schedule models.Schedule
			schedule.Month = month.Month
			exist := models.IsExist(db, "Schedule", bson.M{"month": month.Month})
			if exist {
				helper.ResponseWithJson(w, http.StatusOK,
					helper.Response{Code: 20000, Msg: "Month already exist!"})
				return
			} else {
				if err := adminDao.InsertSchedule(schedule); err != nil {
					helper.ResponseWithJson(w, http.StatusInternalServerError,
						helper.Response{Code: http.StatusInternalServerError, Msg: err.Error()})
					return
				}
				//success
				helper.ResponseWithJson(w, http.StatusOK,
					helper.Response{Code: 20000, Msg: "setMonth success!"})
				return
			}

		}
	} else {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		return
	}

}
