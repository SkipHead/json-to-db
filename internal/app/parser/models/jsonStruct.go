package models

import "time"

type JsonStruct struct {
	Test      string    `json:"test"`
	Test2     int       `json:"test_2"`
	TimeStamp time.Time `json:"time_stamp"`
}
