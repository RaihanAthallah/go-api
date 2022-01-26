package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUser struct {
	IDUser       int    `json:"iduser"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	IDIdentifier int    `json:"id_identifier"`
	Salt         string `json:"salt"`
}

func (HvUser) TableName() string {
	return config.C.Database.Schema.ListClinics + ".hayvee_clinics"
}
