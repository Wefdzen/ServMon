package models

type Server struct {
	Id            uint8
	NameOfService string `json:"nameofservice"`
	Account       string `json:"account"`
	IpServer      string `json:"ipserver"`
	Password      string `json:"password"`
}

type ServerInfo struct {
	IpServer    string
	CoreCount   uint8
	LoadAvg5Min string
	Ram         string
	Memory      string
}
