package users

import (
	"database/sql/driver"
	"os"
)

//go:generate stringer -type=Gender -output=constants_generated.go
type Gender uint8

const (
	Male Gender = iota
	Female
)

//
//func (g Gender) String() string {
//	//TODO implement me
//	return []string{"Male", "Female"}[g]
//}

func (c *Gender) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		switch string(v) {
		case "male":
			*c = Male
			return nil
		case "female":
			*c = Female
			return nil
		}
	}

	return os.ErrInvalid
}

func (c Gender) Value() (driver.Value, error) {
	var db string

	switch c {
	case Male:
		db = "male"
	case Female:
		db = "female"
	default:
		return "", os.ErrInvalid
	}

	return db, nil
}
