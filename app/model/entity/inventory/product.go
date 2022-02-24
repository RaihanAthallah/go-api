package entity

import (
	"github.com/inventory-management-tokobejo/go-api/config"
)

type Product struct {
	Id_Product           int    `gorm:"primary_key:auto_increment;not_null" json:"id_product"`
	SKU                  string `json:"sku"`
	Id_Supplier          int    `json:"id_supplier"`
	Id_Product_Type      int    `json:"id_product_type"`
	Id_Brand             int    `json:"id_brand"`
	Product_Description  string `json:"product_description"`
	Weight               int    `json:"weight"`
	Id_Weight_Type       int    `json:"id_weight_type"`
	Initial_Stock        int    `json:"initial_stock"`
	Id_Location          int    `json:"id_location"`
	Initial_Cost         int    `json:"initial_cost"`
	Buy_Price            int    `json:"buy_price"`
	Wholesale_Price      int    `json:"wholesale_price"`
	Retail_Price         int    `json:"retail_price"`
	Created_Date_Product string `json:"created_date_product"`
	Last_Updated_Product string `json:"last_updated_product"`
}

func (Product) TableName() string {
	return config.C.Database.Schema.Inventory + ".product"
}
