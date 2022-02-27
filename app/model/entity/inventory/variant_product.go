package entity

import (
	"github.com/inventory-management-tokobejo/go-api/config"
)

type Variant struct { 
	Id         		int     `gorm:"primary_key:auto_increment;not_null" json:"id"`
	Id_Product 		int     `json:"id_product"`
	Variant       	string 	`json:"variant"`
	Option_Values   string  `json:"option_values"`
}

func (Variant) TableName() string {
	return config.C.Database.Schema.Inventory + ".variant"
}
