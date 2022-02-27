package response

type ResponseCreatedTracking struct {
	Id         int    `json:"id"`
	Id_Product int    `json:"id_product"`
	Desc       string `json:"desc"`
	Id_Track   int    `json:"id_track"`
}
