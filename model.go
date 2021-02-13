package main
import (
	"database/sql"
	"errors"
	"time"
)

type uptime_entity struct {
	Time time.Time `json:"timestamp"` 
	StatusCode int `json:"statusCode"`
}

func (u *uptime_entity) getUptime(db *sql.DB) error {
	return errors.New("Not implemented")
}