package model

import (
	"database/sql"
	"encoding/json"
)

type NullTime struct {
	sql.NullTime
}

func (v *NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)
}
