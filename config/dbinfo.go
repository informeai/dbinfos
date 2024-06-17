package config
//DBInfoMongoConfigure is struct for configure adapter mongo
type DBInfoMongoConfigure struct {
	Username   string `json:"usename"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}
