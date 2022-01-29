package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUserData struct {
	IDUser       int    `gorm:"primary_key:auto_increment;not_null;column:iduser"`
	IDIdentifier int    `json:"id_identifier"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
}

func (HvUserData) TableName() string {
	return config.C.Database.Schema.User + ".hayvee_user_data"
}