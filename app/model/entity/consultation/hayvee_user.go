package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUser struct {
	IDUser       int    `gorm:"primary_key:auto_increment;not_null;column:iduser"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	IDIdentifier int    `json:"id_identifier"`
	Salt         string `json:"salt"`
}

func (HvUser) TableName() string {
	return config.C.Database.Schema.User + ".hayvee_user"
}
