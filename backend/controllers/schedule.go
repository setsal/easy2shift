package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../helper"
	"../models"
	"github.com/dgrijalva/jwt-go"
)

// Database access object
var (
	dao = models.Schedule{}
)

type UserSchedule struct {
	Month       string   `bson:"time" json:"time"`
	Days        []string `bson:"days" json:"days"`
	Description string   `bson:"desc" json:"desc"`
}

func CreateUserSchedule(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	// decode json
	var data UserSchedule
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// check user schedule exist or not
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

	// get user name and try to insert or update
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var result models.Schedule
		result, err := dao.FindScheduleByMonth(data.Month)
		if err != nil {
			print(err.Error())
		}

		// get record
		var record models.Record = result.List[claims["username"].(string)]
		if len(record.Days) == 0 {
			// create new record for user
			print("no record")
		} else {
			// overwrite the exist value
			var record models.Record
			record.Days = data.Days
			record.Description = data.Description
			if err := dao.UpdateUserRecord(data.Month, claims["username"].(string), record); err != nil {
				helper.ResponseWithJson(w, http.StatusInternalServerError, err.Error())
				return
			} else {
				helper.ResponseWithJson(w, http.StatusOK,
					helper.Response{Code: 20000, Msg: "update success"})
				return
			}
		}
	} else {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		return
	}

	// var schedule models.Schedule
	// var record models.Record

	// record.Days = data.Days
	// record.Description = data.Description

	// schedule.Id = bson.NewObjectId()
	// schedule.Month = data.Time
	// test := map[string]models.Record{
	// 	"admin": record,
	// }
	// schedule.List = test
	// fmt.Println(reflect.TypeOf(schedule.List))

	// if err := dao.InsertSchedule(schedule); err != nil {
	// 	helper.ResponseWithJson(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	helper.ResponseWithJson(w, http.StatusCreated, "create success")
}
