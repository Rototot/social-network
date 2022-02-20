package users

import (
	"database/sql/driver"
	"encoding/json"
	"os"
	"strconv"
)

type UserID int64
type HashedPassword string
type UserInterests []string

// UserID
func (u *UserID) UnmarshalBinary(data []byte) error {
	v, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	*u = UserID(v)

	return nil
}

func (u UserID) MarshalBinary() ([]byte, error) {
	return []byte(strconv.Itoa(int(u))), nil
}

func (u UserID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(u))), nil
}

func (u UserID) String() string {
	return strconv.Itoa(int(u))
}

// UserInterests
func (u UserInterests) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		return json.Unmarshal(v, &u)
	}

	return os.ErrInvalid
}

func (u UserInterests) Value() (driver.Value, error) {
	return json.Marshal(u)
}
