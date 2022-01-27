package response

type ListDoctor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	ServiceTime string `json:"service_time"`
	City        string `json:"city"`
	Contact     string `json:"contact"`
	Province    string `json:"province"`
	PostalCode  string `json:"postal_code"`
	Avatar      string `json:"avatar"`
}
