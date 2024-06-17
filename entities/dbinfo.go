package entities

import (
	"time"
)

// DBInfos is stuct for dbinfos
type DBInfos struct {
	Topic     string    `json:"topic"`
	Info      any       `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
