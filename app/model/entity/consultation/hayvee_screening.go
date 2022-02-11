package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvScreening struct {
	ID         int    `gorm:"primary_key:auto_increment;not_null;column:id"`
	IDUser     int    `gorm:"column:iduser"`
	Number1    string `json:"number1"`
	Number2    string `json:"number2"`
	Number3    string `json:"number3"`
	Number4    string `json:"number4"`
	Number5    string `json:"number5"`
	IsDone     bool   `gorm:"column:isdone"`
	LastUpdate string `json:"last_update"`
}

func (HvScreening) TableName() string {
	return config.C.Database.Schema.Consultation + ".screening"
}
