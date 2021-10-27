package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type IRequest interface {
	Get() []Request
	Save(input *Request) Request
	Update(requestUuid string, input *Request) Request
	Delete(commandUuid string) Request
}

type MRequest struct {
	db *gorm.DB
}

type Request struct {
	Uuid string `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Name string `json:"name"`
}

func (Request)tableName() string {
	return "requests"
}

var AllRequest []Request = []Request{{
	"324e672e-3512-41b0-97dc-065c334f8f7a",
	time.Now(),
	time.Now(),
	"Request 1",
},{
	"6737a85c-3ea4-48ed-bfed-4bdff74b189b",
	time.Now(),
	time.Now(),
	"Request 2",
}}

func NewRequest(_db *gorm.DB) *MRequest {
	output := new(MRequest)
	output.db = _db
	return output
}

func NewDefaultRequest() *MRequest {
	return NewRequest(db)
}

func (r *MRequest) Get() (requests []Request, err error) {
	if findErr := r.db.Find(&requests).Error; findErr != nil {
		err = findErr
		return
	}
	return
}

func (r *MRequest) Save(input *Request) Request {
	return *input
}

func (r *MRequest) Update(requestUuid string, input *Request) Request {
	output := new(Request)
	for _, request := range AllRequest {
		if requestUuid == request.Uuid {
			output.Uuid = request.Uuid
			output.CreateAt = request.CreateAt
			output.Name = input.Name
			output.UpdateAt = time.Now()
			return *output
		}
	}
	return Request{}
}

func (r *MRequest) Delete(commandUuid string) Request {
	fmt.Println("Delete uuid", commandUuid)
	return Request{}
}

