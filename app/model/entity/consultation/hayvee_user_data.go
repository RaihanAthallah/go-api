package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUserData struct {
	IDUser       int    `gorm:"primary_key:auto_increment;not_null;column:iduser"`
	IDIdentifier int    `json:"id_identifier"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Umur         int    `json:"umur"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat       string `json:"alamat" form:"alamat"`
	Kota         string `json:"kota" form:"kota"`
}

func (HvUserData) TableName() string {
	return config.C.Database.Schema.User + ".hayvee_user_data"
}
