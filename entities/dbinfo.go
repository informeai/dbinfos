package entities

// DBInfos is stuct for dbinfos
type DBInfos struct {
	Topic     string `json:"topic"`
	Info      any    `json:"info"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
