package models

import (
	"encoding/json"
	"strings"
	"time"
)

//author: gurbanli

type CustomDate time.Time

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*c = CustomDate(t)
	return nil
}

func (c CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(c)
	return json.Marshal(t)
}
