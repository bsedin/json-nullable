package nullable

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString struct {
	sql.NullString
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullInt64 struct {
	sql.NullInt64
}

type NullBool struct {
	sql.NullBool
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (c *NullString) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(c.String)
}

func (c *NullInt64) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(c.Int64)
}

func (c *NullFloat64) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(c.Float64)
}

func (c *NullBool) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(c.Bool)
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}
	nt.Time, nt.Valid = value.(time.Time), true
	return nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil)
	}
	return nt.Time.MarshalJSON()
}
