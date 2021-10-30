package entity

import (
	"github.com/hayvee-website-development/go-api-hayvee/config"
)

type PubUser struct {
	ID          int    `json:"id"`
	IdMahasiswa int    `json:"id_mahasiswa"`
	Name        string `gorm:"column:name"`
	NIM         int    `json:"nim"`
	Password    string `json:"icon"`
}

func (PubUser) TableName() string {
	return config.C.Database.Schema.ListClinics + ".hayvee_clinics"
}
