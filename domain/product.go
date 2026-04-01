package domain

import (
	"time"
	"strings"
	"database/sql/driver"
	"errors"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	return []byte(`"` + d.In(loc).Format("2006-01-02 15:04:05") + `"`), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *DateTime) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}
	return d.Time, nil
}

func (d *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if t, ok := value.(time.Time); ok {
		d.Time = t
		return nil
	}
	return errors.New("failed to scan DateTime")
}

type Product struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *float64 `json:"price"`
	CreatedAt   *DateTime `json:"created_at"`
	UpdatedAt   *DateTime `json:"updated_at"`
	DeletedAt   *DateTime `json:"deleted_at" gorm:"index"`
}