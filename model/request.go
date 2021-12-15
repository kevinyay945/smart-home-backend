package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type IRequest interface {
	Get() (requests []Request, err error)
	Save(input *Request) (Request, error)
	Update(requestUuid string, input *Request) (Request, error)
	Delete(commandUuid string) error
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

func NewOriginRequest(_db *gorm.DB) IRequest {
	output := new(MRequest)
	output.db = _db
	return output
}

func NewRequest() IRequest {
	return NewOriginRequest(db)
}

func (r *MRequest) Get() (requests []Request, err error) {
	err = r.db.Find(&requests).Error
	return
}

func (r *MRequest) Save(input *Request) (Request, error) {
	err := r.db.Save(input).Error
	return *input, err
}

func (r *MRequest) Update(requestUuid string, input *Request) (Request, error) {
	err := r.db.Model(&Request{}).
		Updates(input).
		Where("uuid = ?", requestUuid).Error
	if err != nil {
		return Request{}, err
	}
	return Request{}, nil
}

func (r *MRequest) Delete(commandUuid string) error {
	fmt.Println("Delete uuid", commandUuid)
	err := r.db.Where("uuid = ?", commandUuid).Delete(&Request{}).Error
	return err
}

