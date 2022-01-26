package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvUserAccess struct {
	ID           int    `json:"id"`
	IDUser       int    `json:"iduser"`
	IDIdentifier int    `json:"id_identifier"`
	Token        string `json:"token"`
	LastUpdate   string `json:"last_update"`
}

func (HvUserAccess) TableName() string {
	return config.C.Database.Schema.ListClinics + ".hayvee_clinics"
}
