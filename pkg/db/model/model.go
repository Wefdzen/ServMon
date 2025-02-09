package model

type RecordAboutServerInfo struct {
	ID          uint
	Time        int64  //Unix time
	NameService string `json:"nameservice"`
	IpServer    string `json:"ipserver"`
	LoadAvg5Min string `json:"loadavg5min"`
	Ram         string `json:"ram"`
	Memory      string `json:"memory"`
}
