package models

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

type Schedule struct {
	Id    bson.ObjectId     `bson:"_id" json:"id"`
	Month string            `bson:"month" json:"month"`
	List  map[string]Record `bson:"record" json:"record"`
}

type Record struct {
	Days        []string `bson:"days" json:"days"`
	Description string   `bson:"desc" json:"desc"`
}

const (
	db         = "Dormnet"
	collection = "Schedule"
)

func (m *Schedule) InsertSchedule(data Schedule) error {
	return Insert(db, collection, data)
}

func (m *Schedule) FindAllSchedule() ([]Schedule, error) {
	var result []Schedule
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (m *Schedule) FindScheduleByMonth(month string) (Schedule, error) {
	var result Schedule
	err := FindOne(db, collection, bson.M{"month": month}, nil, &result)
	return result, err
}

func (m *Schedule) UpdateUserRecord(month string, user string, record Record) error {
	fmt.Println(record)
	data := bson.M{"$set": bson.M{"record." + user: record}}
	return Update(db, collection, bson.M{"month": month}, data)
}
