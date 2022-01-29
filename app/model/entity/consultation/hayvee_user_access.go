package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUserAccess struct {
	ID           int    `gorm:"primary_key:auto_increment;not_null;column:id"`
	IDUser       int    `gorm:"column:iduser"`
	IDIdentifier int    `json:"id_identifier"`
	Token        string `json:"token"`
	LastUpdate   string `json:"last_update"`
}

func (HvUserAccess) TableName() string {
	return config.C.Database.Schema.User + ".hayvee_user_access"
}
