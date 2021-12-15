package schema

import "time"

type Command struct {
	Uuid     string    `json:"uuid" gorm:"primary_key;type:uuid; NOT NULL;default:uuid_generate_v4();"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Url      string    `json:"url"`
}

var AllCommand []Command = []Command{{
	"bbbad37f-c6cd-47f7-907d-add1c4045559",
	time.Now(),
	time.Now(),
	"http://example.com",
}, {
	"bbbad37f-c6cd-47f7-907d-add1c4045558",
	time.Now(),
	time.Now(),
	"http://example2.com",
}}
