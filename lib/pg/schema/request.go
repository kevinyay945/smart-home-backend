package schema

import "time"

type Request struct {
	Uuid     string    `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Name     string    `json:"name" validate:"required"`
}

func (Request) TableName() string {
	return "requests"
}
