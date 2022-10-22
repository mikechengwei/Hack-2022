package entity

type DataSource struct {
	BaseEntity
	Name        string `json:"name"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	StatusPort  string `json:"statusPort"`
	PdAddress   string `json:"pdAddress"`
	Type        int    `json:"type"`
	ServiceName string `json:"serviceName"`
}

func (data *DataSource) TableName() string {
	return "hackathon_datasource"
}
