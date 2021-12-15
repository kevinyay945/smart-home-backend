package model

import (
	"fmt"
	"gorm.io/gorm"
	"smart-home-backend/lib/pg/schema"
)

type IRequest interface {
	Get() (requests []schema.Request, err error)
	Save(input *schema.Request) (schema.Request, error)
	Update(requestUuid string, input *schema.Request) (schema.Request, error)
	Delete(commandUuid string) error
}

type MRequest struct {
	db *gorm.DB
}

func NewOriginRequest(_db *gorm.DB) IRequest {
	output := new(MRequest)
	output.db = _db
	return output
}

func NewRequest() IRequest {
	return NewOriginRequest(db)
}

func (r *MRequest) Get() (requests []schema.Request, err error) {
	err = r.db.Find(&requests).Error
	return
}

func (r *MRequest) Save(input *schema.Request) (schema.Request, error) {
	err := r.db.Save(input).Error
	return *input, err
}

func (r *MRequest) Update(requestUuid string, input *schema.Request) (schema.Request, error) {
	err := r.db.Model(&schema.Request{}).
		Updates(input).
		Where("uuid = ?", requestUuid).Error
	if err != nil {
		return schema.Request{}, err
	}
	return schema.Request{}, nil
}

func (r *MRequest) Delete(commandUuid string) error {
	fmt.Println("Delete uuid", commandUuid)
	err := r.db.Where("uuid = ?", commandUuid).Delete(&schema.Request{}).Error
	return err
}
