package request

type RequestCreatedProduct struct {
	SKU                  string `json:"sku" form:"sku"`
	Id_Supplier          int    `json:"id_supplier" form:"id_supplier"`
	Id_Product_Type      int    `json:"id_product_type" form:"id_product_type"`
	Id_Brand             int    `json:"id_brand" form:"id_brand"`
	Product_Description  string `json:"product_description" form:"product_description"`
	Weight               int    `json:"weight" form:"weight"`
	Id_Weight_Type       int    `json:"id_weight_type" form:"id_weight_type"`
	Initial_Stock        int    `json:"initial_stock" form:"initial_stock"`
	Id_Location          int    `json:"id_location" form:"id_location"`
	Initial_Cost         int    `json:"initial_cost" form:"initial_cost"`
	Buy_Price            int    `json:"buy_price" form:"buy_price"`
	Wholesale_Price      int    `json:"wholesale_price" form:"wholesale_price"`
	Retail_Price         int    `json:"retail_price" form:"retail_price"`
	Created_Date_Product string `json:"created_date_product" form:"created_date_product"`
	Last_Updated_Product string `json:"last_updated_product" form:"last_updated_product"`
}
