package entities

import "time"

type DBInfo struct {
	Topic     string    `json:"topic"`
	Data      any       `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
