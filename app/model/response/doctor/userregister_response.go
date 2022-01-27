package response

type ResultUser struct {
	ID           int    `json:"iduser"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IDIdentifier int    `json:"id_identifier"`
	Salt         string `json:"salt"`
}

type ResponeRegister struct {
	User       interface{} `json:"user"`
	Token      string      `json:"token"`
	JwtToken   string      `json:"jwt_token"`
	ServerTime string      `json:"server_time"`
}
