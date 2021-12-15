package schema

import "time"

type Command struct {
	Uuid     string    `json:"uuid" gorm:"primary_key;type:uuid; NOT NULL;default:uuid_generate_v4();"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Url      string    `json:"url"`
}

func (Command) TableName() string {
	return "commands"
}
