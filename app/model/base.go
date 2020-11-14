package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NullTime struct {
	sql.NullTime
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return json.Marshal(nil)
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	timestamp, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(data), "\""))
	*t = Timestamp(timestamp)
	return err
}

func (t Timestamp) Value() (driver.Value, error) {
	timestamp := time.Time(t)
	return timestamp.Format("2006-01-02 15:04:05"), nil
}
