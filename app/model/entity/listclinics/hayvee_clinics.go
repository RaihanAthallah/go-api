package entity

import "github.com/hayvee-website-development/go-api-hayvee/config"

type HvClinic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	ServiceTime string `json:"service_time"`
	City        string `json:"city"`
	Contact     string `json:"contact"`
	Province    string `json:"province"`
	PostalCode  string `json:"postal_code"`
}

func (HvClinic) TableName() string {
	return config.C.Database.Schema.ListClinics + ".hayvee_clinics"
}
