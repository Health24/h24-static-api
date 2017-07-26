package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/vshaveyko/h24-static-api/config"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}

var DB *gorm.DB

type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func init() {

	fmt.Println(config.DATABASE_URL)

	var err error

	DB, err = gorm.Open("postgres", "user=health24 dbname=health24_production")

	if err != nil {

		fmt.Println(err)

	}

}
