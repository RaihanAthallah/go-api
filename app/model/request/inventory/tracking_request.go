package request

type TrackingRequest struct {
	Id_Product int    `json:"id_product"`
	Desc       string `json:"desc"`
	Id_Track   int    `json:"id_track"`
}
