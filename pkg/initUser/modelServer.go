package inituser

type Server struct {
	Id       uint8
	Account  string `json:"account"`
	IpServer string `json:"ipserver"`
	Password string `json:"password"`
}
