package apis 
type ConnectCampusNetWorkResp struct {
	UserIndex string `json:"userIndex"`
	Result string `json:"result"`
	Message string `json:"message"`
	Forwordurl interface{} `json:"forwordurl"`
	KeepaliveInterval int `json:"keepaliveInterval"`
	ValidCodeURL string `json:"validCodeUrl"`
}