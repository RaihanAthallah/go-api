package response

type ResponseCreatedVariant struct {
	Id            int    `json:"id"`
	Id_Product    int    `json:"id_product"`
	Variant       string `json:"variant"`
	Option_Values string `json:"option_values"`
}
