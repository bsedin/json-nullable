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

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ni.Int64)
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float64)
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nb.Bool)
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
