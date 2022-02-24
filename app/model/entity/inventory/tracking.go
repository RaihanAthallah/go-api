package entity

import (
	"github.com/inventory-management-tokobejo/go-api/config"
)

type Tracking struct {
	Id         int    `gorm:"primary_key:auto_increment;not_null" json:"id"`
	Id_Product int    `json:"id_product"`
	Desc       string `json:"desc"`
	Id_Track   int    `json:"id_track"`
}

func (Tracking) TableName() string {
	return config.C.Database.Schema.Inventory + ".tracking"
}
